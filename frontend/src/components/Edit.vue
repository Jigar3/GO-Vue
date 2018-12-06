<template>
    <div>
        <Cover />

        <router-link to="/"><svg xmlns="http://www.w3.org/2000/svg" id="arrow" width="24" height="24" viewBox="0 0 24 24"><path d="M0 0h24v24H0z" fill="none"/><path d="M20 11H7.83l5.59-5.59L12 4l-8 8 8 8 1.41-1.41L7.83 13H20v-2z"/></svg></router-link>

        <section id="form">
            <form @submit.prevent="processForm">
                <div class="field">
                    <label for="title">Title</label>
                    <input type="text" v-model="data.title" value="data.title">
                </div>

                <div class="field">
                    <label for="content">Content</label>
                    <textarea type="text" rows="8" v-model="data.content"></textarea>
                </div>

                <div class="field">
                    <label for="author">Author</label>
                    <input type="text" v-model="data.author" disabled>
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
    name: "Edit",
    components: {
        Cover
    },
    data() {
        return {
            data: {}
        }
    },
    created() {
        axios.get("http://localhost:8000/view/" + this.$route.params.id)
            .then(resp => {
                this.data = resp.data;
            })
    },
    methods: {
        processForm() {
            axios.patch("http://localhost:8000/update/" + this.$route.params.id , {
                "title": this.data.title,
                "content": this.data.content
            }).then(() => {
                router.push("/");
            })
        }
    }
}
</script>

