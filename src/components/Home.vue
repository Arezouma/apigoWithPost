<template>
  <div class="home">
    <h1>{{ msg }}</h1>
    <div id="box">
      <form class="container" @submit.prevent="processForm">
        <h3>PLEASE ENTER YOUR ARTICLE INFORMATION</h3>
        <p>
          <label for="id">ID</label> <br>
          <input type="id" name="id" id="title" v-model="id" placeholder="Enter the Id number">
        </p>
        <p>
          <label for="title">Title</label> <br>
          <input type="text" name="title" id="title" v-model="title" placeholder="Enter the Title">
        </p>

        <p>
          <label for="author">Author</label><br>
          <input type="text" name="author" id="author" v-model="author" placeholder="Enter the AAuthor">
        </p>
        <p>
          <label for="image">Image</label><br>
          <input type="url" name="image" id="image" v-model="image" placeholder="Enter Image Url">
        </p>
        <div class="button">
          <p>
            <input type="submit" value="Submit" method="POST">  
          </p>
          <p>
            <input type="submit" value="Update" method="PUT">  
          </p>
          <p>
            <input type="submit" value="Delete" method="Delete">  
          </p>
        </div>

      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      msg: 'Welcome to My Api',
      id: '',
      title:'',
      author:'',
      image:''
    }
  },
  methods: {
    processForm: function() {
      id = this.id,
      title = this.title,
      author = this.title,
      image = this.image,
      console.log({Id:this.id,title:this.title,author:this.author,image:this.image})
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container { 
  box-sizing: content-box; 
  margin: auto;
  width: 1000px;
  margin-left: -250px;



}
#box{
  box-sizing: content-box; 
  border: 2px solid green;
  width:500px;
  position: center;
  margin: auto;
  border-radius:25px;
}

h1 {
  text-align: center;
  color: grey;

}
h2 {
  text-align: center;
  color: lightblue;

}
.button p {
  display: inline-block; 
}
input[type=submit] {
    width: 100%;
    background-color: #4CAF50;
    color: white;
    padding: 14px 20px;
    margin: 8px 0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
input {
    width: 30%;
    padding: 12px 20px;
    margin: 8px 0;
    display: inline-block;
    border: 2px solid #ccc;
    border-radius: 4px;
    box-sizing: border-box;
}

</style>
