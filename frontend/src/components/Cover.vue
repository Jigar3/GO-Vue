<template>
  <div class="blog">
    <div class="cover">
      <img src="../assets/cover.jpg" alt id="bgcover">
      <p class="overlay-text">Thoughts in Words</p>
    </div>
  </div>
</template>

<script>
import "reset-css";
import axios from "axios";

export default {
  name: "Cover"
};
</script>

<style scoped>
.cover {
  position: relative;
  text-align: center;
  color: white;
}

#bgcover {
  max-height: 250px;
  width: 100%;
  object-fit: cover;
}

.overlay-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 80px;
  font-family: "Righteous", cursive;
}
</style>
