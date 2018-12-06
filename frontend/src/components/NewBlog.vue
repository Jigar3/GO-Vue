<template>
    <div>
        <Cover />

        <router-link to="/"><svg xmlns="http://www.w3.org/2000/svg" id="arrow" width="24" height="24" viewBox="0 0 24 24"><path d="M0 0h24v24H0z" fill="none"/><path d="M20 11H7.83l5.59-5.59L12 4l-8 8 8 8 1.41-1.41L7.83 13H20v-2z"/></svg></router-link>
        
        <section id="form">
            <form @submit.prevent="processForm">
                <div class="field">
                    <label for="title">Title</label>
                    <input type="text" v-model="title">
                </div>

                <div class="field">
                    <label for="content">Content</label>
                    <textarea type="text" rows="8" v-model="content"></textarea>
                </div>

                <div class="field">
                    <label for="author">Author</label>
                    <input type="text" v-model="author">
                </div>
                <button type="submit" id="save">SAVE</button>
            </form>
        </section>
    </div>    
</template>

<script>
import Cover from "./Cover";
import axios from "axios";
import router from "../router.js"

export default {
    name: "NewBlog",
    components: {
        Cover
    },
    data() {
        return {
            title: "",
            content: "",
            author: ""
        }
    },
    methods: {
        processForm() {
            axios.post("http://localhost:8000/new", {
                "title": this.title,
                "content": this.content,
                "author": this.author
            }).then(() => {
                router.push("/");
            })
        }
    }
}
</script>

<style>
#arrow {
    position: relative;
    left: calc(5%);
    padding-top: calc(10px);
    height: 60px;
    width: 60px;
}

.field {
    display: flex;
    margin: 0 auto;
    max-width: 30%;
    justify-content: space-between;
    padding: calc(1%);
}

.field input  {
    border: 1px black solid;
    width: calc(70%);
    height: 25px;
}

.field textarea {
    border: 1px black solid;
    width: calc(70%);
}

.field label {
    font-size: 18px;
    font-style: "PT Sans", sans-serif;
}

#form {
    display: inline;
}

#save {
    margin: 0 auto;
    position: relative;
    left: calc(54%);
    margin: 20px;
    padding: 5px 15px;
    background: #B1DDC9;
    color: white;
    font-style: "PT Sans", sans-serif;
    border: none;
    cursor: pointer;
}

</style>


