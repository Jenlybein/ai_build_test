import React, { useState, useMemo, useEffect } from "react";
import { Layout, Row, Col } from "tea-component";
import FilterPanel from "./FilterPanel";
import FieldSelector from "./FieldSelector";
import LogTable from "./LogTable";
import Message from "./Message";
import { getLogs } from "../services/requestApi";
import { DEFAULT_SELECTED_FIELDS, STORAGE_KEYS } from "../constants";
import { getStorage, setStorage } from "../utils/storageUtil";

const LogQuery = () => {
  // 预选的字段
  const [selectedFields, setSelectedFields] = useState(
    getStorage(STORAGE_KEYS.SELECTED_FIELDS) || DEFAULT_SELECTED_FIELDS
  );
  // 日志数据项
  const [logData, setLogData] = useState([]);
  // 加载状态
  const [loading, setLoading] = useState(false);
  // 消息
  const [message, setMessage] = useState(null);
  // 记录总数
  const [recordCount, setRecordCount] = useState(1);
  // 分页数据
  const [pageData, setPageData] = useState({
    pageSize: 10,
    pageIndex: 1,
  });
  // 查询参数
  const [queryParams, setQueryParams] = useState({
    app_id: "",
    platform: "all",
    timeRange: [],
    user_id: "",
    entrance_time: "",
    id: "",
  });

  // 监听字段选择变化并保存
  useEffect(() => {
    setStorage(STORAGE_KEYS.SELECTED_FIELDS, selectedFields);
  }, [selectedFields]);

  // 计算查询参数
  const computedParams = useMemo(() => {
    const params = {
      ...queryParams,
      start_time: queryParams.timeRange[0]
        ? queryParams.timeRange[0].format("YYYY-MM-DD HH:mm:ss")
        : "",
      end_time: queryParams.timeRange[1]
        ? queryParams.timeRange[1].format("YYYY-MM-DD HH:mm:ss")
        : "",
      page_index: pageData.pageIndex,
      page_size: pageData.pageSize,
    };
    delete params.timeRange;
    return params;
  }, [queryParams, pageData]);

  const handleSearch = async () => {
    try {
      setLoading(true);
      setMessage(null); // 清除之前的消息

      const response = await getLogs(computedParams);

      setLogData(response.data);
      setRecordCount(response.total);
      setMessage({
        content: response.message,
        type: "success",
      });
    } catch (error) {
      setMessage({
        content: error.message,
        type: error.type,
      });
      setLogData([]);
      setRecordCount(0);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Layout>
      <Layout.Content>
        <div style={{ padding: "20px" }}>
          {message && (
            <Message
              content={message.content}
              type={message.type}
              onClose={() => setMessage(null)}
              style={{ marginBottom: "20px" }}
            />
          )}
          <FilterPanel
            onSearch={handleSearch}
            queryParams={queryParams}
            setQueryParams={setQueryParams}
          />

          <Row style={{ marginTop: "20px" }}>
            <Col span={4}>
              <FieldSelector
                selectedFields={selectedFields}
                onFieldsChange={setSelectedFields}
              />
            </Col>
            <Col span={20} style={{ paddingLeft: "20px" }}>
              <LogTable
                selectedFields={selectedFields}
                data={logData}
                loading={loading}
                recordCount={recordCount}
                setPageData={setPageData}
                onSearch={handleSearch}
              />
            </Col>
          </Row>
        </div>
      </Layout.Content>
    </Layout>
  );
};

export default LogQuery;
