<template>
  <div class="bible-search-container">
    <!-- 搜索区域 -->
    <div class="search-section">
      <div class="search-box">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="输入关键词或经文(如: 创1:1 或 爱)"
          @keyup.enter="searchBible"
        />
        <button @click="searchBible">搜索</button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div class="loading" v-if="loading">搜索中...</div>

    <!-- 结果区域 -->
    <div class="results-section" v-if="results.length > 0">
      <div class="verse-card" v-for="(verse, index) in results" :key="index">
        <h3 class="verse-reference">
          {{ verse.reference }}
        </h3>
        <p class="verse-text" v-html="highlightKeywords(verse.text)"></p>
      </div>
    </div>

    <div class="no-results" v-if="searched && results.length === 0 && !loading">
      未找到匹配的经文
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
      searched: false,
      loading: false,
    };
  },
  methods: {
    async searchBible() {
      if (!this.searchQuery.trim()) return;

      this.searched = true;
      this.loading = true;
      this.results = [];

      try {
        const response = await fetch(
          `/lmrl/api/search?q=${encodeURIComponent(this.searchQuery)}`
        );
        const data = await response.json();
        this.results = data.results || [];
      } catch (error) {
        console.error("搜索失败:", error);
        this.results = [];
      } finally {
        this.loading = false;
      }
    },

    isVerseReference(query) {
      // 简单的经文引用检测（如"创1:1"）
      return /[a-zA-Z\u4e00-\u9fa5]+\d*:\d+/.test(query);
    },

    highlightKeywords(text) {
      if (!this.searchQuery || this.isVerseReference(this.searchQuery)) {
        return text;
      }

      const query = this.searchQuery.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
      const regex = new RegExp(query, "gi");
      return text.replace(
        regex,
        (match) => `<span class="highlight">${match}</span>`
      );
    },
  },
};
</script>

<style scoped>
.bible-search-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: "Noto Sans SC", sans-serif;
}

.search-section {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.search-box {
  display: flex;
  margin-bottom: 10px;
}

.search-box input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px 0 0 4px;
  font-size: 16px;
}

.search-box button {
  padding: 10px 15px;
  background: #4a6fa5;
  color: white;
  border: none;
  border-radius: 0 4px 4px 0;
  cursor: pointer;
  font-size: 16px;
}

.advanced-options {
  display: flex;
  gap: 10px;
  margin-bottom: 10px;
}

.advanced-options input,
.advanced-options select {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.toggle-advanced {
  background: none;
  border: none;
  color: #4a6fa5;
  cursor: pointer;
  font-size: 14px;
  padding: 5px 0;
}

.results-section {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.verse-card {
  background: white;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.verse-reference {
  margin: 0 0 10px 0;
  color: #333;
  font-size: 18px;
}

.verse-text {
  margin: 0;
  line-height: 1.6;
  font-size: 16px;
}

.highlight {
  background-color: #fff3b0;
  font-weight: bold;
}

.no-results,
.loading {
  text-align: center;
  padding: 20px;
  color: #666;
}

/* 响应式设计 */
@media (max-width: 600px) {
  .search-box {
    flex-direction: column;
  }

  .search-box input {
    border-radius: 4px;
    margin-bottom: 5px;
  }

  .search-box button {
    border-radius: 4px;
  }

  .advanced-options {
    flex-direction: column;
  }
}
</style>
