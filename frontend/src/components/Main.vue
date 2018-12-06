<template>
  <div class="blog">
    <Cover />

    <section class="container">
      <div class="card" v-for="blog in content" :key="blog.id">
        <h4>
          <router-link to="/blog">{{ blog.title }}</router-link>
        </h4>
        <h2>{{blog.author}} <span style="font-weight: 500" >|</span> {{ blog.time.slice(0,16) }}</h2>
        <p>{{ blog.content }}</p>
      </div>
    </section>

    <router-link to="/new">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
        <path d="M0 0h24v24H0z" fill="none"></path>
        <path
          d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm5 11h-4v4h-2v-4H7v-2h4V7h2v4h4v2z"
        ></path>
      </svg>
    </router-link>
  </div>
</template>

<script>
import "reset-css";
import axios from "axios";
import Cover from "./Cover";

export default {
  name: "Main",
  data() {
    return {
      content: []
    };
  },
  created() {
    axios.get("http://localhost:8000/view").then(data => {
      return this.content = data.data;
    })
  },
  components: {
    Cover
  }
};
</script>

<style scoped>

.container {
  margin: 0 auto;
  width: 65%;
  margin-top: calc(5%);
}

h4,
h4 a {
  font-family: "Open Sans Condensed", sans-serif;
  font-size: 28px;
  margin-bottom: 20px;
  cursor: pointer;
  text-decoration: none;
  color: black;
}

h2 {
  font-family: "Muli", sans-serif;
  font-size: 18px;
  margin-bottom: 20px;
}

p {
  font-family: "PT Sans", sans-serif;
  font-size: 16px;
  margin-bottom: calc(8%);
}

svg {
  position: fixed;
  top: calc(85%);
  left: calc(90%);
  height: calc(10%);
  width: calc(10%);
  fill: #79cac1;
  cursor: pointer;
}
</style>
