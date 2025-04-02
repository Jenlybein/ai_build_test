// 获取存储的值
export const getStorage = (key) => {
  try {
    const value = localStorage.getItem(key);
    return value ? JSON.parse(value) : null;
  } catch (error) {
    console.error("获取存储失败:", error);
    return null;
  }
};

// 保存值到存储
export const setStorage = (key, value) => {
  try {
    localStorage.setItem(key, JSON.stringify(value));
  } catch (error) {
    console.error("保存到存储失败:", error);
  }
};

// 从存储中删除值
export const removeStorage = (key) => {
  try {
    localStorage.removeItem(key);
  } catch (error) {
    console.error("从存储中删除失败:", error);
  }
};
