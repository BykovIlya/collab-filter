<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Table of products</h4>
            </b-col>
          </b-row>
        </template>
        <b-table id="products"
                 striped
                 show-empty
                 :items="items"
                 :fields="fields"
                 :current-page="currentPage"
                 :per-page="perPage"
                 :total-rows="totalRows"
                 :busy.sync="isBusy"
                 ref="table"
        >
        </b-table>
        <b-row>
          <b-col sm="12">
            <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"/>
          </b-col>
        </b-row>
      </b-card>
    </b-container>
  </div>
</template>

<script>
  export default {
    name: 'Products',
    data () {
      return {
        formUrl: 'http://localhost:5001',
        fields: [
          {
            key: 'itemid',
          },
          {
            key: 'count',
          },
        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy:false,
        totalRows:1,
        currentPage: 1,
        perPage: 15,
        fileProducts:null,
        items:[]
      }
    },
    created(){
    this.getItems()
  },
    methods:{
      getItems(ctx){
        let url = "http://localhost:5001/products";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
           // if(result.body.length > 0) {
              this.items = result.body;
              this.totalRows = this.items.length;
              return result.body
          //  }
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
