import * as XLSX from "xlsx";

/**
 * 将数据导出为Excel文件
 * @param {Array} data - 要导出的数据数组
 * @param {Array} columns - 列配置数组
 * @param {string} filename - 导出文件名
 */
export const exportToExcel = (data, columns, filename = "export.xlsx") => {
  try {
    // 准备Excel数据
    const excelData = data.map((item) => {
      const row = {};
      columns.forEach((col) => {
        row[col.title] = item[col.key];
      });
      return row;
    });

    // 创建工作簿
    const wb = XLSX.utils.book_new();
    const ws = XLSX.utils.json_to_sheet(excelData);

    // 设置列宽
    const colWidths = columns.map((col) => ({
      wch: Math.max(col.title.length, 10), // 根据标题长度设置最小宽度
    }));
    ws["!cols"] = colWidths;

    // 将工作表添加到工作簿
    XLSX.utils.book_append_sheet(wb, ws, "Sheet1");

    // 导出文件
    XLSX.writeFile(wb, filename);
  } catch (error) {
    console.error("导出Excel失败:", error);
    throw new Error("导出Excel失败");
  }
};
