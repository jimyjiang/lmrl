<template>
  <div class="google-style-search">
    <!-- 搜索区域 - Google风格 -->
    <div class="search-container">
      <div class="search-box">
        <div class="search-icon">
          <svg viewBox="0 0 24 24" width="20" height="20">
            <path
              fill="#9AA0A6"
              d="M15.5 14h-.79l-.28-.27a6.5 6.5 0 0 0 1.48-5.34c-.47-2.78-2.79-5-5.59-5.34a6.505 6.505 0 0 0-7.27 7.27c.34 2.8 2.56 5.12 5.34 5.59a6.5 6.5 0 0 0 5.34-1.48l.27.28v.79l4.25 4.25c.41.41 1.08.41 1.49 0 .41-.41.41-1.08 0-1.49L15.5 14zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"
            />
          </svg>
        </div>
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索圣经经文或关键词"
          @keyup.enter="handleSearch"
          class="search-input"
          ref="searchInput"
        />
        <div class="search-actions" v-if="searchQuery">
          <button @click="clearSearch" class="clear-button">
            <svg viewBox="0 0 24 24" width="24" height="24">
              <path
                fill="#70757A"
                d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
              />
            </svg>
          </button>
          <div class="divider"></div>
          <button @click="handleSearch" class="search-button">
            <svg viewBox="0 0 24 24" width="24" height="24">
              <path
                fill="#4285F4"
                d="M15.5 14h-.79l-.28-.27a6.5 6.5 0 0 0 1.48-5.34c-.47-2.78-2.79-5-5.59-5.34a6.505 6.505 0 0 0-7.27 7.27c.34 2.8 2.56 5.12 5.34 5.59a6.5 6.5 0 0 0 5.34-1.48l.27.28v.79l4.25 4.25c.41.41 1.08.41 1.49 0 .41-.41.41-1.08 0-1.49L15.5 14zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div class="loading-indicator" v-if="loading">
      <div class="loading-bar"></div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 历史记录 -->
      <div class="history-panel">
        <h4>搜索历史</h4>
        <ul class="history-list">
          <li
            v-for="(item, index) in history"
            :key="index"
            class="history-item"
          >
            <span @click="searchFromHistory(item)" class="history-text">{{ item }}</span>
            <button @click.stop="deleteHistoryItem(index)" class="delete-history-btn">×</button>
          </li>
        </ul>
      </div>

      <!-- 结果区域 -->
      <div class="results-panel" :class="{ fullscreen: isFullScreen }">
        <div class="results-header">
          <div class="result-stats" v-if="hasResults">
            约 {{ results.length }} 条结果
          </div>
          <button @click="toggleFullScreen" class="fullscreen-button">
            <span v-if="isFullScreen">退出全屏</span>
            <span v-else>全屏显示</span>
            <span class="shortcut-tip">
              ({{ isMac ? "Cmd" : "Alt" }} + Enter)</span
            >
          </button>
        </div>

        <div class="search-results" v-if="hasResults">
          <div
            class="result-item"
            v-for="(verse, index) in results"
            :key="`${verse.reference}-${index}`"
          >
            <div class="verse-reference">
              <a href="#" class="verse-link">{{ verse.reference }}</a
              >&nbsp;&nbsp;
              <span
                class="verse-text"
                v-html="highlightText(verse.text)"
              ></span>
            </div>
          </div>
        </div>

        <!-- 无结果提示 -->
        <div class="no-results" v-if="showEmptyState">
          <div class="no-results-icon">
            <svg viewBox="0 0 24 24" width="48" height="48">
              <path
                fill="#9AA0A6"
                d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8zm.5-13H11v6l5.2 3.2.8-1.3-4.5-2.7z"
              />
            </svg>
          </div>
          <h3 class="no-results-title">
            没有找到与"{{ searchQuery }}"匹配的经文
          </h3>
          <p class="no-results-tips">
            建议：
            <br />• 检查输入是否正确 <br />• 尝试其他关键词 <br />•
            使用完整的经文引用（如"创世记1:1"）
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "BibleSearch",
  data() {
    return {
      searchQuery: "",
      results: [],
      loading: false,
      searched: false,
      history: [],
      isFullScreen: false,
    };
  },
  computed: {
    isSearchDisabled() {
      return !this.searchQuery.trim();
    },
    hasResults() {
      return this.results.length > 0;
    },
    showEmptyState() {
      return this.searched && !this.hasResults && !this.loading;
    },
    isMac() {
      return navigator.platform.toUpperCase().indexOf("MAC") >= 0;
    },
  },
  mounted() {
    this.loadHistory();
    window.addEventListener("keydown", this.handleKeyDown);
  },
  beforeUnmount() {
    window.removeEventListener("keydown", this.handleKeyDown);
  },
  methods: {
    async handleSearch() {
      if (this.isSearchDisabled || this.loading) return;

      const queryToSearch = this.searchQuery.trim();
      this.searched = true;
      this.loading = true;
      this.results = [];

      try {
        const query = encodeURIComponent(queryToSearch);
        const response = await fetch(`/lmrl/api/search?q=${query}`);

        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        this.results = data.results || [];
        if (this.results.length > 0) {
          this.updateHistory(queryToSearch);
        }
      } catch (error) {
        console.error("搜索失败:", error);
        this.results = [];
      } finally {
        this.loading = false;
      }
    },

    clearSearch() {
      this.searchQuery = "";
      this.results = [];
      this.searched = false;
      this.$refs.searchInput.focus();
    },

    isVerseReference(text) {
      return /^[\u4e00-\u9fa5a-zA-Z]+\s*\d+[\s:]\s*\d+$/.test(text.trim());
    },

    highlightText(text) {
      if (!this.searchQuery || this.isVerseReference(this.searchQuery)) {
        return text;
      }

      const query = this.searchQuery.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
      const regex = new RegExp(query, "gi");
      return text.replace(
        regex,
        (match) => `<span class=\"highlight\">${match}</span>`
      );
    },

    loadHistory() {
      const savedHistory = localStorage.getItem("bible_search_history");
      if (savedHistory) {
        this.history = JSON.parse(savedHistory);
      }
    },

    updateHistory(query) {
      if (!query) return;
      // 移除已存在的相同记录，并添加到最前面
      const index = this.history.indexOf(query);
      if (index > -1) {
        this.history.splice(index, 1);
      }
      this.history.unshift(query);
      // 限制历史记录数量
      if (this.history.length > 20) {
        this.history.pop();
      }
      this.saveHistory();
    },

    deleteHistoryItem(index) {
      this.history.splice(index, 1);
      this.saveHistory();
    },

    saveHistory() {
      localStorage.setItem(
        "bible_search_history",
        JSON.stringify(this.history)
      );
    },

    searchFromHistory(query) {
      this.searchQuery = query;
      this.handleSearch();
    },

    toggleFullScreen() {
      this.isFullScreen = !this.isFullScreen;
    },

    handleKeyDown(event) {
      const isMac = this.isMac;
      if ((isMac ? event.metaKey : event.altKey) && event.key === "Enter") {
        event.preventDefault();
        this.toggleFullScreen();
      }
    },
  },
};
</script>

<style scoped>
/* Google风格设计 */
.google-style-search {
  max-width: 100%;
  margin: 0 auto;
  padding: 20px;
  font-family: "Roboto", "Noto Sans SC", sans-serif;
  color: #202124;
}

/* 搜索框 */
.search-container {
  position: relative;
  margin-bottom: 20px;
  max-width: 652px;
  margin-left: auto;
  margin-right: auto;
}

.search-box {
  display: flex;
  align-items: center;
  border: 1px solid #dfe1e5;
  border-radius: 24px;
  padding: 8px 16px;
  height: 44px;
  transition: all 0.3s;
}

.search-box:hover {
  box-shadow: 0 1px 6px rgba(32, 33, 36, 0.1);
}

.search-box:focus-within {
  box-shadow: 0 1px 6px rgba(32, 33, 36, 0.2);
}

.search-icon {
  margin-right: 12px;
  display: flex;
  align-items: center;
}

.search-input {
  flex: 1;
  border: none;
  outline: none;
  font-size: 16px;
  background: transparent;
  color: #202124;
}

.search-actions {
  display: flex;
  align-items: center;
  margin-left: 8px;
}

.clear-button {
  background: none;
  border: none;
  padding: 4px;
  cursor: pointer;
  opacity: 0.7;
  transition: opacity 0.2s;
}

.clear-button:hover {
  opacity: 1;
}

.divider {
  width: 1px;
  height: 24px;
  background: #dfe1e5;
  margin: 0 8px;
}

.search-button {
  background: none;
  border: none;
  padding: 4px;
  cursor: pointer;
}

/* 加载指示器 */
.loading-indicator {
  height: 3px;
  overflow: hidden;
  margin-bottom: 20px;
  max-width: 652px;
  margin-left: auto;
  margin-right: auto;
}

.loading-bar {
  height: 100%;
  width: 100%;
  background: linear-gradient(90deg, #4285f4, #34a853, #fbbc05, #ea4335);
  animation: loading 2s linear infinite;
  transform-origin: left;
  transform: scaleX(0);
}

@keyframes loading {
  0% {
    transform: scaleX(0);
  }
  50% {
    transform: scaleX(1);
  }
  100% {
    transform: scaleX(0) translateX(100%);
  }
}

/* 主内容区 */
.main-content {
  display: flex;
  gap: 20px;
  box-sizing: border-box;
  flex-grow: 1; /* Allow main content to grow and fill available space */
  min-height: 0; /* Allow main content to shrink if needed */
  align-items: flex-start; /* Align items to the top */
  flex-direction: row; /* Arrange children horizontally */
}
.history-panel {
  flex-grow: 0;
  flex-shrink: 0;
  flex-basis: 200px; /* Explicitly set width for flex item */
  min-width: 200px; /* Ensure it doesn't shrink below this */
  border-right: 1px solid #e0e0e0;
  padding-right: 20px;
  box-sizing: border-box;
  overflow-y: auto; /* Always show scrollbar if content overflows vertically */
  height: 100%; /* Fill the height of main-content */
}

.history-panel h4 {
  margin-top: 0;
  font-size: 16px;
  color: #70757a;
}

.history-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.history-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px;
  cursor: pointer;
  border-radius: 4px;
  font-size: 14px;
  color: #4d5156;
}

.history-item:hover {
  background-color: #f1f3f4;
}

.history-text {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  flex-grow: 1;
}

.delete-history-btn {
  background: none;
  border: none;
  color: #9aa0a6;
  font-size: 16px;
  cursor: pointer;
  padding: 0 4px;
  margin-left: 8px;
  visibility: hidden;
}

.history-item:hover .delete-history-btn {
  visibility: visible;
}

.delete-history-btn:hover {
  color: #202124;
}

.results-panel {
  flex-grow: 1; /* Allow it to grow */
  flex-shrink: 1; /* Allow it to shrink */
  flex-basis: 0; /* Allow it to shrink to 0 before content */
  min-width: 0; /* Important for flex items to shrink */
  box-sizing: border-box;
  overflow-x: hidden; /* Prevent horizontal scrollbar */
  overflow-y: auto; /* Always show scrollbar if content overflows vertically */
  height: 100%; /* Fill the height of main-content */
}
.results-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}
.fullscreen-button {
  background: none;
  border: 1px solid #dfe1e5;
  border-radius: 4px;
  padding: 4px 8px;
  font-size: 12px;
  color: #70757a;
  cursor: pointer;
}
.fullscreen-button:hover {
  background-color: #f1f3f4;
}
.shortcut-tip {
  color: #9aa0a6;
}
.results-panel.fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: #fff;
  z-index: 1000;
  overflow-y: auto;
  padding: 20px;
  box-sizing: border-box;
}
.results-panel.fullscreen .history-panel {
  display: none;
} /* 搜索结果 */
.search-results {
  padding-top: 6px;
}
.result-stats {
  color: #70757a;
  font-size: 14px;
}
.result-item {
  margin-bottom: 6px;
  max-width: 600px;
}
.verse-reference {
  margin-bottom: 4px;
}
.verse-link {
  color: #1a0dab;
  text-decoration: none;
  font-size: 20px;
  line-height: 1.3;
  font-weight: normal;
}
.verse-link:hover {
  text-decoration: underline;
}
.verse-text {
  color: #4d5156;
  line-height: 1.58;
  font-size: 14px;
}
.highlight {
  font-weight: bold;
  color: #202124;
  background-color: rgba(66, 133, 244, 0.1);
} /* 无结果 */
.no-results {
  text-align: center;
  padding: 80px 0;
}
.no-results-icon {
  margin-bottom: 24px;
}
.no-results-title {
  font-size: 20px;
  font-weight: normal;
  color: #202124;
  margin: 0 0 16px;
}
.no-results-tips {
  color: #70757a;
  font-size: 14px;
  line-height: 1.6;
  margin: 0;
  text-align: left;
  display: inline-block;
} /* 响应式设计 */
@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
    height: auto; /* Allow height to be auto in column layout */
  }
  .history-panel {
    flex-basis: auto;
    border-right: none;
    border-bottom: 1px solid #e0e0e0;
    padding-right: 0;
    padding-bottom: 10px;
    margin-bottom: 10px;
    height: auto; /* Allow height to be auto in column layout */
  }
  .results-panel {
    height: auto; /* Allow height to be auto in column layout */
  }
}
@media (max-width: 600px) {
  .google-style-search {
    padding: 16px;
  }
  .search-box {
    height: 40px;
    padding: 6px 12px;
  }
  .verse-link {
    font-size: 18px;
  }
  .verse-text {
    font-size: 13px;
  }
}
</style>
