export const formatDate = () => {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, "0"); // 月份从 0 开始，需要 +1
    const day = String(now.getDate()).padStart(2, "0");
    return `${year}-${month}-${day}`;
};