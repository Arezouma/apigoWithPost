<template>
	<div id="Nav">
	  <router-link to="/" class="nav">HOME</router-link>
      <router-link to="articles" class="nav">ARTICLES</router-link>
      <router-link to="contact" class="nav">CONTACT</router-link>
		
	</div>
</template>

<script>
	export default {
		name: 'Nav'
	}
	
</script>

<style>

	.nav { 
		text-decoration: none;
		margin: 10px;
		color:black;


	 }

	#Nav{
		text-align: center;
        margin: 20px;
        font-size: 15px;
    }
</style>