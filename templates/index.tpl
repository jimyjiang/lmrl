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
  <link href="https://cdn.jsdelivr.net/npm/font-awesome@4.7.0/css/font-awesome.min.css" rel="stylesheet">
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
          <a href="/灵命日粮/{{.Filename}}"  target="_blank" class="text-primary hover:text-secondary transition-custom font-medium flex items-center">
            <i class="fa fa-angle-right text-primary/70 mr-2"></i>
            <span class="truncate">{{.Title}}</span>
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
  </main>

</body>
</html>