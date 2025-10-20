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
  font-family: Avenir, Helvetica, Arial, sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

header {
  text-align: center;
  margin-bottom: 30px;
}

header h1 {
  color: #2c3e50;
}
</style>
