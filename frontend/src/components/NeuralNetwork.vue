<template>
 <div>
   <b-container fluid>
     <b-card border-variant="">
       <template slot="header">
         <b-row>
           <b-col sm="5">
             <h4 class="card-title">Рекомендация для нового покупателя</h4>
           </b-col>
         </b-row>
       </template>
       <div>
         <h4 class="card-title" align="left">Введите данные:</h4>
         <b-form inline>
           <b-input class="mb-2 mr-sm-2 mb-sm-0" id="gender" placeholder="пол (0-муж, 1-жен)" v-model="gender"/>
           <b-input-group left="@" class="mb-2 mr-sm-2 mb-sm-0">
             <b-input id="age" placeholder="Возраст" v-model="age"/>
           </b-input-group>
           <b-input-group left="@" class="mb-2 mr-sm-2 mb-sm-0">
             <b-input id="category" placeholder="Категория товара (от 1 до 10)" v-model="category"/>
           </b-input-group>
           <b-input-group left="@" class="mb-2 mr-sm-2 mb-sm-0">
             <b-input id="price" placeholder="Цена товара" v-model="price"/>
           </b-input-group>
           <b-button variant="primary" id="see" @click="seeNN()">Посмотреть</b-button>
         </b-form>
         <br/>
         <br/>
         <h4 id="#answer1" class="card-title" align="left"> {{ answer }}</h4>
         <br/>
         <br/>
       </div>
       <div>
         <h4 class="card-title" align="left">Введите данные:</h4>
         <b-form inline>
           <b-input class="mb-2 mr-sm-2 mb-sm-0" id="iduser" placeholder="id пользователя" v-model="iduser"/>
           <b-input-group left="@" class="mb-2 mr-sm-2 mb-sm-0">
             <b-input id="idproduct" placeholder="id продукта" v-model="idproduct"/>
           </b-input-group>
           <b-button variant="primary" id="see" @click="seeNN()">Посмотреть</b-button>
         </b-form>
         <br/>
         <br/>
         <h4 id="#answer2" class="card-title" align="left"> {{ answer2 }}</h4>
       </div>
     </b-card>
   </b-container>
 </div>
</template>

<script>
  import $ from 'jquery';
  export default {
        name: "NeuralNetwork",
        data() {
          return {
            formUrl: 'http://localhost:5001',
            gender:'',
            age:'',
            category:'',
            price:'',
            answer:'',
            iduser:'',
            idproduct:'',
            answer2:''
          }
        },
        methods: {
          seeNN: function () {
            this.answer = '';
            /* let age = $('#age').text();
             let category = $('#category').text();
             let gender = $('#gender').text();
             let price = $('#price').text();
             */
            let age = new String(this.age)
            let category = new String(this.category)
            let gender = new String(this.gender)
            let price = new String(this.price)
            let url = 'http://localhost:5001/neuralnetwork/' + age + '/' + gender + '/' + category + '/' + price;
            this.$http.get(url).then(result => {
              console.log(result);
              if (result === "1") {
                this.answer = "Данный продукт рекомендован данному пользователю"
              } else {
                this.answer = "Данный продукт не рекомендован данному пользователю"
              }
            })
          },

        }
    }
</script>

<style scoped>

</style>
