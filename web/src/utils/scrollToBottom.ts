// DOM 引用，指定类型为 HTMLDivElement | null
const logContainer = ref<HTMLDivElement | null>(null);
    // 滚动状态
const isAtBottom = ref(true); // 是否滚动条在底部

    // 方法：滚动到最底部
export const scrollToBottom = () => {
  nextTick(() => {
    if (isAtBottom.value && logContainer.value) {
      logContainer.value.scrollTop = logContainer.value.scrollHeight;
    }
  });
};

    // 滚动事件处理
export const handleScroll = () => {
  if (!logContainer.value) return;   
    const { scrollTop, scrollHeight, clientHeight } = logContainer.value;
    const atBottom = scrollTop + clientHeight >= scrollHeight - 10; // 判断是否接近底部
      
    isAtBottom.value = atBottom; // 更新状态
  };
