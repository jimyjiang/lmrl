<template>
  <div id="app">
    <header>
      <h1>圣经查经工具</h1>
    </header>
    <main>
      <BibleSearch />
    </main>
  </div>
</template>

<script>
import BibleSearch from "./components/BibleSearch.vue";

export default {
  name: "App",
  components: {
    BibleSearch,
  },
};
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    "Noto Sans SC", sans-serif;
  max-width: 1280px;
  margin: 0 auto;
  padding: 16px;
  text-align: left;
}

header {
  margin-bottom: 20px;
  padding-bottom: 12px;
  border-bottom: 1px solid #e0e0e0;
}

header h1 {
  color: #202124;
  font-size: 1.5rem;
  font-weight: 500;
  margin: 0;
  padding: 0;
  text-align: center;
}

main {
  padding: 0 8px;
}

/* 响应式调整 */
@media (max-width: 600px) {
  #app {
    padding: 12px;
  }

  header {
    margin-bottom: 16px;
  }

  header h1 {
    font-size: 1.3rem;
  }
}
</style>
