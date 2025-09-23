<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <link rel="icon" type="image/png" href="/favicon-96x96.png" sizes="96x96" />
  <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
  <link rel="shortcut icon" href="/favicon.ico" />
  <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
  <link rel="manifest" href="/site.webmanifest" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>《旷野吗哪》灵修信息</title>
  <script src="https://cdn.tailwindcss.com"></script>
  <script src="assets/javascript/tailwindcss-3.4.17.js"></script>
  <!-- <link href="https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css" rel="stylesheet"> -->
  <link href="assets/css/font-awesome.min.css" rel="stylesheet">
  <script>
    tailwind.config = {
      theme: {
        extend: {
          colors: {
            primary: '#E67E22',
            secondary: '#D35400',
            neutral: '#ECF0F1',
            dark: '#2C3E50'
          },
          fontFamily: {
            sans: ['Inter', 'system-ui', 'sans-serif'],
          },
        },
      }
    }
  </script>
  <style type="text/tailwindcss">
    @layer utilities {
      .text-shadow-sm {
        text-shadow: 0 1px 2px rgba(0,0,0,0.1);
      }
      .transition-custom {
        transition: all 0.3s ease;
      }
    }
    /* 自定义进度条样式 */
#progressBar::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 16px;
  height: 16px;
  background: #E67E22;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.2s;
}

#progressBar::-webkit-slider-thumb:hover {
  transform: scale(1.2);
  background: #D35400;
}

#progressBar::-moz-range-thumb {
  width: 16px;
  height: 16px;
  background: #E67E22;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.2s;
}

#progressBar::-moz-range-thumb:hover {
  transform: scale(1.2);
  background: #D35400;
}
  </style>
</head>
<body class="bg-gray-50 font-sans text-gray-800 min-h-screen">
  <!-- 顶部导航 -->
  <header class="bg-white shadow-sm sticky top-0 z-10 transition-custom">
    <div class="container mx-auto px-4 py-4 flex justify-between items-center">
      <h1 class="text-[clamp(1.5rem,3vw,2rem)] font-bold text-dark">
        <i class="fa fa-book-open text-primary mr-2"></i>《旷野吗哪》灵修信息
      </h1>
      <div class="hidden md:flex space-x-4">
        <a href="#" class="bg-primary text-white px-4 py-2 rounded-md hover:bg-secondary transition-custom">
          <i class="fa fa-home mr-1"></i>首页
        </a>
      </div>
      <button class="md:hidden text-primary text-xl">
        <i class="fa fa-bars"></i>
      </button>
    </div>
  </header>

  <!-- 主要内容 -->
  <main class="container mx-auto px-4 py-8">
    <!-- 介绍部分 -->
    <section class="bg-white rounded-xl shadow-md p-6 mb-8 transform hover:shadow-lg transition-custom">
      <div class="flex flex-wrap items-center gap-3 mb-4">
        <span class="bg-primary/10 text-primary px-3 py-1 rounded-full text-sm font-medium">栽培</span>
        <span class="bg-primary/10 text-primary px-3 py-1 rounded-full text-sm font-medium">旷野吗哪</span>
      </div>
      
      <div class="mb-4">
        <h2 class="text-lg font-semibold text-dark mb-2 flex items-center">
          <i class="fa fa-user-circle-o text-primary mr-2"></i>讲师: 孙大中
        </h2>
        <p class="text-gray-700 leading-relaxed">
          节目主持孙大中，每天送上一段经文，一篇灵修短文，一句金句，一篇精致短讲和祷告，陪伴你每天灵修。
        </p>
      </div>
    </section>

    <!-- 灵修列表 -->
    <section class="bg-white rounded-xl shadow-md overflow-hidden">
      <div class="bg-primary text-white p-4">
        <h2 class="text-xl font-semibold flex items-center">
          <i class="fa fa-calendar-check-o mr-2"></i>灵修日程
        </h2>
      </div>
      
      <div class="divide-y divide-gray-100">
        <!-- 灵修项目 - 使用循环和条件样式 -->
      {{range .SermonList}}
      <div class="p-4 hover:bg-gray-50 transition-custom flex flex-col md:flex-row md:items-center justify-between">
        <div class="flex-1 min-w-0">
          <a href="/灵命日粮/{{.Filename}}" class="text-primary hover:text-secondary transition-custom font-medium flex items-center">
            <i class="fa fa-angle-right text-primary/70 mr-2"></i>
            <span class="truncate">{{.Title}}</span>
            <button onclick="copyToClipboard(event, '{{.Title}}', this)" class="ml-2 text-gray-500 hover:text-gray-700" title="复制标题">
              <i class="fa fa-copy"></i>
            </button>
            <span id="copyFeedback" class="hidden text-sm text-green-500 ml-2">已复制!</span>
          </a>
          
          <div class="flex flex-wrap gap-x-4 gap-y-1 mt-2 text-sm text-gray-600">
            <span class="flex items-center">
              <i class="fa fa-user text-gray-400 mr-1"></i>{{.Speaker}}
            </span>
            <span class="flex items-center">
              <i class="fa fa-calendar text-gray-400 mr-1"></i>{{.Date}}
            </span>
            <span class="flex items-center">
              <i class="fa fa-clock-o text-gray-400 mr-1"></i>{{.Duration}}
            </span>
            <span class="flex items-center">
              <i class="fa fa-file-text-o text-gray-400 mr-1"></i>{{.FileSize}}
            </span>
          </div>
        </div>
        <span class="mt-2 md:mt-0 text-gray-500 text-sm bg-gray-50 px-3 py-1 rounded-full whitespace-nowrap">{{.Date}}</span>
      </div>
      {{end}}
      </div>
    </section>
<!-- 音频播放弹出层 -->
<div id="audioPlayerModal" class="fixed inset-0 z-50 hidden items-center justify-center bg-black bg-opacity-70">
  <div class="relative bg-white rounded-xl shadow-2xl w-full max-w-md mx-4 overflow-hidden">
    <!-- 关闭按钮 -->
    <button id="closePlayer" class="absolute top-4 right-4 text-gray-500 hover:text-gray-700 transition-colors">
      <i class="fa fa-times text-2xl"></i>
    </button>
    
    <!-- 播放器主体 -->
    <div class="p-6">
      <!-- 音频信息 -->
      <div class="flex items-center mb-6">
        <div class="bg-primary/10 p-3 rounded-lg mr-4">
          <i class="fa fa-music text-primary text-2xl"></i>
        </div>
        <div>
          <h3 id="audioTitle" class="font-semibold text-lg text-gray-800 truncate max-w-[250px]">讲道标题</h3>
          <p id="audioSpeaker" class="text-gray-500 text-sm">孙大中</p>
        </div>
      </div>
      
      <!-- 音频控制 -->
      <audio id="modalAudioPlayer" class="w-full"></audio>
      
      <!-- 自定义控制条 -->
      <div class="mt-4 space-y-3">
        <!-- 进度条 -->
        <div class="flex items-center space-x-3">
          <span id="currentTime" class="text-xs text-gray-500 w-10">0:00</span>
          <input type="range" id="progressBar" class="flex-1 h-2 bg-gray-200 rounded-full appearance-none cursor-pointer" min="0" max="100" value="0">
          <span id="duration" class="text-xs text-gray-500 w-10">0:00</span>
        </div>
        
        <!-- 控制按钮 -->
        <div class="flex justify-center items-center space-x-6">
          <button id="rewindBtn" class="text-gray-600 hover:text-primary transition-colors">
            <i class="fa fa-step-backward text-xl"></i>
          </button>
          <button id="playPauseBtn" class="bg-primary text-white rounded-full w-12 h-12 flex items-center justify-center hover:bg-secondary transition-colors">
            <i class="fa fa-play text-xl"></i>
          </button>
          <button id="forwardBtn" class="text-gray-600 hover:text-primary transition-colors">
            <i class="fa fa-step-forward text-xl"></i>
          </button>
        </div>
      </div>
      <div class="mt-6">
        <h4 class="text-sm font-medium text-gray-700 mb-2">快速跳转</h4>
        <div class="flex flex-wrap gap-2">
          <!-- 示例时间点 -->
          <button onclick="jumpToTime(60)" class="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full text-sm transition-colors">
            1:00 - 引言
          </button>
          <button onclick="jumpToTime(135)" class="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full text-sm transition-colors">
            2:15 - 诗歌
          </button>
          <button onclick="jumpToTime(320)" class="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full text-sm transition-colors">
            5:20 - 读经
          </button>
          <button onclick="jumpToTime(440)" class="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full text-sm transition-colors">
            7:20 - 经文讲解
          </button>
          <button onclick="jumpToTime(1640)" class="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full text-sm transition-colors">
            27:20 - 祷告
          </button>
        </div>
      </div>
    </div>
  </div>
</div>
  </main>
<script lang="javascript">
  // 获取DOM元素
const audioPlayerModal = document.getElementById('audioPlayerModal');
const modalAudioPlayer = document.getElementById('modalAudioPlayer');
const closePlayer = document.getElementById('closePlayer');
const playPauseBtn = document.getElementById('playPauseBtn');
const progressBar = document.getElementById('progressBar');
const currentTimeEl = document.getElementById('currentTime');
const durationEl = document.getElementById('duration');
const audioTitle = document.getElementById('audioTitle');
const audioSpeaker = document.getElementById('audioSpeaker');

// 初始化播放器状态
let isPlaying = false;

// 为所有音频链接添加点击事件
document.querySelectorAll('a[href$=".mp3"]').forEach(link => {
  link.addEventListener('click', function(e) {
    e.preventDefault(); // 阻止默认跳转
    
    // 设置音频信息
    const title = this.dataset.title || this.textContent.trim();
    const speaker = this.dataset.speaker || "孙大中";
    
    audioTitle.textContent = title;
    audioSpeaker.textContent = speaker;
    
    // 设置音频源
    modalAudioPlayer.src = this.href;
    
    // 显示弹出层
    audioPlayerModal.classList.remove('hidden');
    audioPlayerModal.classList.add('flex');
    
    // 自动播放（注意：浏览器可能会阻止自动播放）
    modalAudioPlayer.play().then(() => {
      isPlaying = true;
      playPauseBtn.innerHTML = '<i class="fa fa-pause text-xl"></i>';
    }).catch(error => {
      console.log("自动播放被阻止:", error);
    });
  });
});

// 关闭播放器
closePlayer.addEventListener('click', () => {
  audioPlayerModal.classList.add('hidden');
  audioPlayerModal.classList.remove('flex');
  modalAudioPlayer.pause();
  isPlaying = false;
});

// 播放/暂停控制
playPauseBtn.addEventListener('click', () => {
  if (isPlaying) {
    modalAudioPlayer.pause();
    playPauseBtn.innerHTML = '<i class="fa fa-play text-xl"></i>';
  } else {
    modalAudioPlayer.play();
    playPauseBtn.innerHTML = '<i class="fa fa-pause text-xl"></i>';
  }
  isPlaying = !isPlaying;
});

// 进度条更新
modalAudioPlayer.addEventListener('timeupdate', () => {
  const currentTime = modalAudioPlayer.currentTime;
  const duration = modalAudioPlayer.duration;
  
  // 更新进度条
  if (!isNaN(duration)) {
    const progressPercent = (currentTime / duration) * 100;
    progressBar.value = progressPercent;
    
    // 更新时间显示
    currentTimeEl.textContent = formatTime(currentTime);
    durationEl.textContent = formatTime(duration);
  }
});

// 拖动进度条
progressBar.addEventListener('input', () => {
  const seekTime = (progressBar.value / 100) * modalAudioPlayer.duration;
  modalAudioPlayer.currentTime = seekTime;
});

// 音频结束处理
modalAudioPlayer.addEventListener('ended', () => {
  isPlaying = false;
  playPauseBtn.innerHTML = '<i class="fa fa-play text-xl"></i>';
});

// 快进/快退按钮
document.getElementById('rewindBtn').addEventListener('click', () => {
  modalAudioPlayer.currentTime = Math.max(0, modalAudioPlayer.currentTime - 15);
});

document.getElementById('forwardBtn').addEventListener('click', () => {
  modalAudioPlayer.currentTime = Math.min(modalAudioPlayer.duration, modalAudioPlayer.currentTime + 15);
});

// 跳转到指定时间(秒)
function jumpToTime(seconds) {
  const audio = document.getElementById('modalAudioPlayer');
  audio.currentTime = seconds;
  
  // 视觉反馈
  const button = event.target;
  button.classList.add('bg-primary', 'text-white');
  setTimeout(() => {
    button.classList.remove('bg-primary', 'text-white');
  }, 1000);
}

// 格式化时间显示 (秒 → MM:SS)
function formatTime(seconds) {
  const mins = Math.floor(seconds / 60);
  const secs = Math.floor(seconds % 60);
  return `${mins}:${secs < 10 ? '0' : ''}${secs}`;
}

function copyToClipboard(event, text, button) {
    // 方法1：尝试使用现代API
    if (navigator.clipboard) {
        navigator.clipboard.writeText(text).then(function() {
            showFeedback(button);
            event.preventDefault(); // 阻止默认行为
            event.stopPropagation(); // 阻止冒泡
        }).catch(function(err) {
            fallbackCopy(text, button);
            event.preventDefault(); // 阻止默认行为
            event.stopPropagation(); // 阻止冒泡
        });
    } 
    // 方法2：使用老式方法作为后备
    else {
        fallbackCopy(text, button);
        event.preventDefault(); // 阻止默认行为
        event.stopPropagation(); // 阻止冒泡
    }
}

function fallbackCopy(text, button) {
    // 创建临时textarea元素
    const textarea = document.createElement('textarea');
    textarea.value = text;
    textarea.style.position = 'fixed';  // 避免滚动到底部
    document.body.appendChild(textarea);
    textarea.select();
    
    try {
        // 尝试执行复制命令
        const successful = document.execCommand('copy');
        if (successful) {
            showFeedback(button);
        } else {
            console.error('复制失败');
        }
    } catch (err) {
        console.error('复制失败:', err);
    }
    
    // 移除临时元素
    document.body.removeChild(textarea);
}

function showFeedback(button) {
    const feedback = button.nextElementSibling;
    feedback.classList.remove('hidden');
    setTimeout(() => feedback.classList.add('hidden'), 2000);
}
</script>
</body>
</html>