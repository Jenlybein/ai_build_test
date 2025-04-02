import React, { useState } from "react";
import { Table, Text, StatusTip, Pagination, Button } from "tea-component";
import { ALL_FIELDS } from "../constants";
import { exportToExcel } from "../utils/exportUtil";
const { expandable, sortable } = Table.addons;

const LogTable = ({
  selectedFields,
  data,
  loading,
  recordCount,
  setPageData,
  onSearch,
}) => {
  const [expandedKeys, setExpandedKeys] = useState([]);

  // 换页
  const handlePagingChange = (pageData) => {
    setPageData({ ...pageData });
    onSearch();
  };

  // 导出Excel
  const handleExport = () => {
    const columns = selectedFields.map((field) => ({
      key: field,
      title: ALL_FIELDS.find((f) => f.key === field).label || field,
    }));
    exportToExcel(data, columns, "日志数据.xlsx");
  };

  // 生成唯一的记录key
  const getRecordKey = (record) => {
    return JSON.stringify(record);
  };

  // 未展开获取信息
  const getColumns = () => {
    return selectedFields.map((field) => ({
      key: field,
      header: ALL_FIELDS.find((f) => f.key === field).label || field,
      render: (record) => {
        const value = record[field];
        if (!value) {
          return <Text theme="weak">-</Text>;
        } else if (value instanceof Date) {
          return <Text>{value.toLocaleString()}</Text>;
        }
        if (typeof value === "object") {
          return <Text>{JSON.stringify(value)}</Text>;
        }
        return <Text>{String(value)}</Text>;
      },
    }));
  };

  // 展开逻辑
  const renderExpandedContent = (record) => {
    // 获取所有字段的值
    const allFieldValues = ALL_FIELDS.map((field) => ({
      key: field.key,
      label: field.label,
      value: record[field.key],
    })).filter((item) => item.value != null);

    return (
      <Table
        records={allFieldValues}
        recordKey="key"
        columns={[
          { key: "label", header: "字段", width: 150 },
          { key: "value", header: "值", render: (item) => String(item.value) },
        ]}
      />
    );
  };

  // 当前排序列
  const [sorts, setSorts] = useState([]);
  return (
    <Table
      columns={getColumns()}
      records={[...data].sort(sortable.comparer(sorts))}
      recordKey={getRecordKey}
      bordered
      style={{ height: "100vh", overflowY: "scroll" }}
      topTip={
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
          }}
        >
          <Pagination
            recordCount={recordCount}
            pageSizeOptions={[10, 20, 50, 100]}
            onPagingChange={handlePagingChange}
          />
          <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
            <Button type="primary" onClick={handleExport}>
              导出Excel
            </Button>
            <StatusTip status={loading ? "loading" : "none"} />
          </div>
        </div>
      }
      addons={[
        sortable({
          columns: [
            "data_time",
            "app_id",
            "user_id",
            "entrance_id",
            "stamp",
            "entrance_time",
            "id",
          ],
          value: sorts,
          onChange: (value) => setSorts(value),
        }),
        expandable({
          expandedKeys,
          onExpandedKeysChange: (keys, { event }) => {
            if (event) {
              event.stopPropagation();
            }
            setExpandedKeys(keys);
          },
          render: renderExpandedContent,
          gapCell: 1,
        }),
      ]}
    />
  );
};

export default LogTable;
