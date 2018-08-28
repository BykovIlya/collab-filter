<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Table of users</h4>
            </b-col>
          </b-row>
        </template>
        <b-table id="users"
                 striped
                 show-empty
                 :items="getItems"
                 :fields="fields"
                 ref="table"
        >
        </b-table>
      </b-card>
    </b-container>
  </div>
</template>

<script>
  export default {
    name: 'Users',
    data () {
      return {
        formUrl: 'http://localhost:5000',
        fields: [
          {
            label: 'Visitor',
            key: 'visitorid_string',
          },
          {
            label: 'Count',
            key: 'items',
          },
          {
          items:[{
              label: 'Item',
              key: 'items.itemid_string',
            },

          ]
          }

        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy: false,
        totalRows:0,
        fileProducts:null
      }
    },
    methods:{
      getItems(ctx){
        let url = "http://localhost:5000/users";
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
