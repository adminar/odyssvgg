<template>
  <div class="page-404">
    <div class="content">
      <SvgIcon name="内容加载失败" :size="500" />
      <div class="prompt">
        <div class="title">404</div>
        <div class="text">抱歉，访问的页面不存在~</div>
        <a-button type="primary" @click="onBack">立即返回</a-button>
      </div>
    </div>
    <div class="game-container" @keydown.space.prevent="jump" tabindex="0">
      <div class="dino" :style="{ bottom: dinoY + 'px', left: dinoX + 'px', width: '80px', height: '80px' }"></div>
      <div
        class="obstacle"
        v-for="(obstacle, index) in obstacles"
        :key="index"
        :style="{ left: obstacle.x + 'px', height: obstacle.height + 'px', width: '60px' }"
      ></div>
      <div class="game-over" v-if="isGameOver">
        <div>Game Over</div>
        <div>跳过的障碍物总数: {{ skippedObstacles }}</div>
        <button @click="restartGame">重新开始</button>
      </div>
      <button v-if="!isGameOver && !gameStarted" @click="startGame" class="start-button">开一局</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { loadingPage } from "@/utils/loading-page";
import { useRouter } from "vue-router";

const router = useRouter();
const onBack = () => {
  router.push({ path: '/home' });
};

const dinoY = ref(0);
const dinoX = ref(0); // 存储 Dino 的水平位置
const jumpHeight = 80; // 跳跃高度
const maxObstacleHeight = jumpHeight * 0.6; // 障碍物最大高度为跳跃高度的80%
const obstacles = ref([]);
const isGameOver = ref(false);
const skippedObstacles = ref(0); // 跳过的障碍物计数
const gameStarted = ref(false); // 游戏是否开始
let obstacleTimer = null;
let lastTime = 0;

const jump = () => {
  if (dinoY.value === 0) {
    dinoY.value = jumpHeight;
    setTimeout(() => {
      dinoY.value = 0;
    }, 300); // 300ms 让跳跃更自然
  }
};

const generateObstacles = () => {
  const newObstacle = {
    x: window.innerWidth,
    height: Math.random() * (maxObstacleHeight - 20) + 20 // 障碍物高度在20到maxObstacleHeight之间
  };
  obstacles.value.push(newObstacle);
};

const moveObstacles = () => {
  const deltaTime = (performance.now() - lastTime) / 1000; // 计算时间差
  obstacles.value.forEach(obstacle => {
    obstacle.x -= 400 * deltaTime; // 每秒移动400个像素
  });

  // 检查并移除越界的障碍物，并更新跳过的障碍物计数
  const initialLength = obstacles.value.length;
  obstacles.value = obstacles.value.filter(obstacle => obstacle.x >= -10); // 保留在视口内的障碍物
  skippedObstacles.value += initialLength - obstacles.value.length; // 增加跳过的障碍物计数

  // 检测碰撞
  obstacles.value.forEach(obstacle => {
    // 检查 Dino 的水平位置与障碍物的位置
    const dinoLeft = dinoX.value;
    const dinoRight = dinoX.value + 30; // Dino 的宽度
    const obstacleLeft = obstacle.x;
    const obstacleRight = obstacle.x + 10; // 障碍物的宽度

    if (obstacleRight > dinoLeft && obstacleLeft < dinoRight && dinoY.value === 0) {
      isGameOver.value = true;
      clearInterval(obstacleTimer);
    }
  });

  lastTime = performance.now(); // 更新最后时间
};

const startGame = () => {
  isGameOver.value = false;
  gameStarted.value = true; // 标记游戏已经开始
  obstacles.value = [];
  dinoY.value = 0;
  dinoX.value = (window.innerWidth - 30) / 2 - 150; // 将Dino的位置设置在中间
  lastTime = performance.now();
  skippedObstacles.value = 0; // 重置跳过的障碍物计数

  obstacleTimer = setInterval(generateObstacles, Math.random() * 2000 + 1000); // 随机间隔生成障碍物
  requestAnimationFrame(gameLoop); // 启动游戏循环
};

const restartGame = () => {
  startGame(); // 重新开始游戏
};

const gameLoop = () => {
  if (!isGameOver.value) {
    moveObstacles();
    requestAnimationFrame(gameLoop); // 继续请求下一帧
  }
};

onMounted(() => {
  loadingPage.done(200);
  // 这里可以考虑不启动游戏，直到用户点击“开始游戏”
});

onBeforeUnmount(() => {
  clearInterval(obstacleTimer);
});
</script>

<style lang="scss" scoped>
.page-404 {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 40px 0; // 上下留白，让内容不完全贴边框

  .content {
    display: flex;
    align-items: center;
    padding: 20px;
    margin-bottom: 20px; // 与下方游戏区域留出间距
    .prompt {
      margin-left: 20px;
      .title {
        font-size: 80px;
        color: $color-text-1;
      }
      .text {
        font-size: $font-size-body-3;
        color: $color-text-2;
        margin-bottom: 10px;
      }
    }
  }

  .game-container {
    width: 100%;
    height: 300px;
    position: relative;
    background: #ffffff;
    overflow: hidden;
    outline: none;
  }

  .dino {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 80px;
    height: 80px;
    background-image: url('car.png');
    background-size: contain;
    background-repeat: no-repeat;
    transition: bottom 0.3s ease;
  }

  .obstacle {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 60px;
    background-image: url('no.png');
    background-size: contain;
    background-repeat: no-repeat;
  }

  .game-over {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 30px;
    color: red;
  }

  .start-button {
    position: absolute;
    font-size: 20px;
    padding: 10px 20px;
    cursor: pointer;
    // Center the start button in the game container
    left: 50%;
    transform: translateX(-50%);
    bottom: 10%;  // Adjust based on preferred vertical placement
  }
}
</style>
