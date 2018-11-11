<template>
	<div class="side">
		<h1>Please Enter your article Info</h1>
		<form id="app" @submit="checkForm" action="/something" method="post">
  
		  <p v-if="errors.length">
		    <b>Please correct the following error(s):</b>
		    <ul>
		      <li v-for="error in errors">{{ error }}</li>
		    </ul>
		  </p>
		  
		  <p>
		    <label for="title">Title:</label>
		    <input type="text" name="title" id="title" v-model="title">
		  </p>

		  <p>
		    <label for="author">Author:</label>
		    <input type="text" name="author" id="author" v-model="author">
		  </p>
		  <p>
		    <label for="image">Image:</label>
		    <input type="url" name="image" id="image" v-model="image">
		  </p>

		  <p>
		    <input type="submit" value="Submit">  
		  </p>

		</form>
	</div>
</template>

<script>
	export default{
		name: "Side",

		data:{
		  	errors:[],
		    title:null,
		    author:null,
		    image:null
		  },
		  methods:{
		    checkForm:function(e) {
		      if(this.title && this.author && this.image) return true;
		      this.errors = [];
		      if(!this.title) this.errors.push("Title required.");
		      if(!this.author) this.errors.push("Author required.");
		      if(!this.image) this.errors.push("Image required.");
		      e.preventDefault();
		    }
		  }
		}
	
</script>

<style scoped>
	input {
  		margin-left: 10px;
	}
	.side {
		float: right;
		width: 20%;
	}
	
</style>