<template>
  <div class="hello">
    <h1>Create Your Article</h1>
    <div v-for="article in articles">
         <h2>{{article.Title}}</h2> 
         <img :src="article.Image">
         <h3>Author: {{article.Author}}</h3>
         
    </div>
    <Side></Side>
  </div>
</template>

<script>
  import Side from './Side'
  export default {
    name: 'Articles',
    components: {Side},
    data () {
      return {
        articles: []
      }
    },
    created: function(){
      fetch('http://localhost:8000/data')
      .then(r => r.json())
      .then(json => {
          this.articles = json;
      });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
img {
  object-fit: cover;
  height: 300px;
  width: 25%;
}
h1 {
  text-align: center;
  color: grey;

}
h2 {
  text-align: center;
  color: lightblue;

}
.hello {
  float:left;
    width: 70%;
}
</style>