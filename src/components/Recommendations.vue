<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Recommendations</h4>
            </b-col>
            <div class ="float-right">
              <b-btn v-b-modal.modalPrevent>Enter the user</b-btn>
            </div>
          </b-row>
        </template>
        <b-table id="recommendations"
                 striped
                 show-empty
                 :items="getItems"
                 :fields="fields"
                 ref="table"
        >
        </b-table>
      </b-card>
    </b-container>
    <b-modal id="modalPrevent"
             ref="modal"
             title="Enter the user"
             @ok="handleOk"
             @shown="clearName">
      <form @submit.stop.prevent="handleSubmit">
        <b-form-input type="text"
                      placeholder="Enter the user"
                      v-model="nameOfRecommendation"></b-form-input>
      </form>
    </b-modal>
  </div>
</template>

<script>
  export default {
    name: 'Recommendations',
    data () {
      return {
        formUrl: 'http://localhost:5000',
        fields: [
          {
            label: 'Item',
            key: 'product',
          },
          {
            label: 'Score',
            key: 'mpRating',
          },
        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy: false,
        totalRows:0,
        fileProducts:null,
        nameOfRecommendation: ''
      }
    },
    methods:{
      clearName () {
        this.nameOfRecommendation = ''
      },
      handleOk (evt) {
        evt.preventDefault()
        if (!this.nameOfRecommendation) {
          alert('Please enter the user')
        } else {
          this.handleSubmit()
        }
      },
      handleSubmit () {
        let myVisitor = new String(this.nameOfRecommendation)
        let url = "http://localhost:5000/recommendations/personal"
        this.$http.post(url, myVisitor, null).then(result => {
          console.log(result);
          if (result.status === 200) {
            this.$refs.table.refresh();
          }
        },error =>{
          console.log(error);
        });
        this.clearName()
        this.$refs.modal.hide()
      },
      getItems(ctx){
        let url = "http://localhost:5000/recommendations";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              return result.body
            }
          }
          return []
        },error =>{
          this.isBusy = false;
          console.log("ERROR",error);
        });
      },
      delete: function (url, data, callback) {
        return this.$http.delete(url,data,null).then(result => {
          callback(result.body);
        },error =>{
          console.log("ERROR",error);
          if (error.status === 422){
            callback(error.body);
          }
          return
        });
      },
      post: function (url, data, callback) {
        console.log(data)
        return this.$http.post(url,data,null).then(result => {
          callback(result);
        },error =>{
          callback(error);
          return
        });
      },
      put: function (url, data, callback) {
        return this.$http.put(url,data,null).then(result => {
          callback(result);
        },error =>{
          callback(error);
          return
        });
      },

    },
  }

</script>
