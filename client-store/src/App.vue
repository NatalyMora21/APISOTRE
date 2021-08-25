<template>

  <div>
    <table class="table">
      <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Name</th>
          <th scope="col">Age</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="buyer in buyers" :key="buyer.id"> 

          <td>{{buyer.id}}</td>
          <td>{{buyer.name}}</td>
          <td>{{buyer.age}}</td>
        
        </tr>
      </tbody>
    </table>
  
  </div>

</template>

<script>

  import axios from 'axios'


  export default {
    name: 'App',
    data(){
      return {
        buyers: null
      }
    },
    mounted(){
      this.getAllBuyers();
      
    },
    methods: {
     async getAllBuyers (){
        console.log("Todos Buyers")
        try{
          this.buyers= await (await axios.get("http://localhost:4000/")).data;
          console.log(this.buyers.data)      

        }catch(err){
          console.log(err)
        }

      }
    }

  }
</script>

<style>

</style>
