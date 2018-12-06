<template>
  <div class="blog">
    <Cover />

    <router-link to="/"><svg xmlns="http://www.w3.org/2000/svg" id="arrow" width="24" height="24" viewBox="0 0 24 24"><path d="M0 0h24v24H0z" fill="none"/><path d="M20 11H7.83l5.59-5.59L12 4l-8 8 8 8 1.41-1.41L7.83 13H20v-2z"/></svg></router-link>

    <section class="container">
      <div class="card">
        <h4>
          {{ data.title }}
        </h4>
      <div id="line">
          <h2>{{data.author}} <span style="font-weight: 500" >|</span> {{ data.time.slice(0,16) }}</h2>
        <div>
          <router-link v-bind:to="{ name: 'update', params: { id: data.id }}">
            <svg id="edit" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M3 17.25V21h3.75L17.81 9.94l-3.75-3.75L3 17.25zM20.71 7.04c.39-.39.39-1.02 0-1.41l-2.34-2.34c-.39-.39-1.02-.39-1.41 0l-1.83 1.83 3.75 3.75 1.83-1.83z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
          </router-link>
          <svg v-on:click="remove" id="delete" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/><path d="M0 0h24v24H0z" fill="none"/></svg>
          </div>
      </div>
        <p>{{ data.content }}</p>
      </div>
    </section>

    </div>
</template>

<script>
import Cover from "../components/Cover";
import axios from "axios";
import router from "../router.js"

export default {
  name: "Blog",
  data() {
    return {  
      data: {}
    }
  },
  components: {
    Cover
  },
  created() {
    axios.get("http://localhost:8000/view/" + this.$route.params.id)
      .then(resp => {
        this.data = resp.data
      })
  },
  methods: {
    remove() {
      axios.delete("http://localhost:8000/delete/" + this.$route.params.id)
        .then(() => {
          router.push("/");
        })
    }
  } 

}
</script>

<style scoped>
.container {
  margin: 0 auto;
  width: 65%;
  margin-top: calc(0%);
}

h4, h4 a {
  font-family: "Open Sans Condensed", sans-serif;
  font-size: 34px;
  margin-bottom: 30px;
  cursor: pointer;
  text-decoration: none;
  color: black;
}

h2 {
  font-family: "Muli", sans-serif;
  font-size: 24px;
  margin-bottom: 30px;
}

p {
  font-family: "PT Sans", sans-serif;
  font-size: 20px;
  margin-bottom: calc(8%);
}

#delete, #edit {
  fill: #D25D42;
  height: 40px;
  width: 40px;
  padding-bottom: 15px;
  box-sizing: border-box;
  cursor: pointer;
  opacity: 0.7;
}

#line {
  display: flex;
  justify-content: space-between;
}

</style>


