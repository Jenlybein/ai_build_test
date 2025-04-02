package repository

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/models"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/gin-gonic/gin"
)

type Repository struct {
	conn driver.Conn
}
func printHex(s string) {
	fmt.Printf("%s (hex): %x\n", s, s)
}
func InitDB() (*Repository, error){
	config := config.LoadClickHouseConfig()
	ip := config.ClickHouseIP
	username := config.ClickHouseUsername
	password := config.ClickHousePassword
	database := config.ClickHouseDatabase

	log.Println("正在连接数据库...")
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{ip},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		Debug: true,
		Settings: map[string]interface{}{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ClickHouse: %v", err)
	}

	return &Repository{conn: conn}, nil
}


func (r *Repository) GetLogs(c *gin.Context) {
	var params models.LogQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		log.Printf("参数解析错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  fmt.Sprintf("参数解析错误: %v", err),
			"params": c.Request.URL.Query(),
		})
		return
	}

	log.Printf("解析后的查询参数: %+v\n", params)

	// 构建基础查询条件
	baseQuery := "FROM kv_7 WHERE 1=1"
	args := []interface{}{}

	if params.StartTime != "" {
		baseQuery += " AND data_time >= ?"
		args = append(args, params.StartTime)
	}
	if params.EndTime != "" {
		baseQuery += " AND data_time <= ?"
		args = append(args, params.EndTime)
	}
	if params.UserID != "" {
		baseQuery += " AND user_id LIKE ?"
		args = append(args, "%"+params.UserID+"%")
	}
	if params.AppID != "" && params.AppID != "all" {
		baseQuery += " AND app_id LIKE ?"
		args = append(args, "%"+params.AppID+"%")
	}
	if params.Platform != "" && params.Platform != "all" {
		baseQuery += " AND platform = ?"
		args = append(args, params.Platform)
	}
	if params.EntranceTime != "" {
		baseQuery += " AND entrance_time = ?"
		args = append(args, params.EntranceTime)
	}
	if params.ID != "" {
		baseQuery += " AND id LIKE ?"
		args = append(args, "%"+params.ID+"%")
	}

	// 获取总记录数
	countQuery := "SELECT COUNT(*) " + baseQuery
	var total uint64
	err := r.conn.QueryRow(context.Background(), countQuery, args...).Scan(&total)
	if err != nil {
		log.Printf("获取总记录数错误: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("获取总记录数错误: %v", err),
		})
		return
	}

	// 构建分页查询
	query := "SELECT * " + baseQuery + " ORDER BY data_time ASC LIMIT ? OFFSET ?"
	args = append(args, params.PageSize, (params.PageIndex-1)*params.PageSize)

	log.Printf("最终SQL查询: %s\n", query)
	log.Printf("查询参数: %v\n", args)

	// 执行查询
	ctx := context.Background()
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		log.Printf("查询执行错误: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("查询执行错误: %v", err),
			"query": query,
			"args":  args,
		})
		return
	}
	defer rows.Close()

	var logs []models.KV7
	for rows.Next() {
		var k models.KV7
		if err := rows.ScanStruct(&k); err != nil {
			log.Printf("数据解析错误: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       fmt.Sprintf("数据解析错误: %v", err),
				"current_row": fmt.Sprintf("%+v", k),
			})
			return
		}
		logs = append(logs, k)
	}

	if err := rows.Err(); err != nil {
		log.Printf("结果遍历错误: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":           fmt.Sprintf("结果遍历错误: %v", err),
			"processed_count": len(logs),
		})
		return
	}

	log.Printf("查询成功完成，返回 %d 条记录\n", len(logs))

	if len(logs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "未找到匹配的记录",
			"data":    []models.KV7{},
			"total":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("成功获取 %d 条记录", len(logs)),
		"data":    logs,
		"total":   total,
	})
}
