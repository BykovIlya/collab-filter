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
                 :items="items"
                 :fields="fields"
                 :current-page="currentPage"
                 :per-page="perPage"
                 :total-rows="totalRows"
                 :busy.sync="isBusy"
                 ref="table"
        >
          <template slot="show_details" slot-scope="row">
            <!-- we use @click.stop here to prevent emitting of a 'row-clicked' event  -->
            <b-button size="sm" @click.stop="row.toggleDetails" class="mr-2">
              {{ row.detailsShowing ? 'Hide' : 'Show'}} Details
            </b-button>
            <!-- In some circumstances you may need to use @click.native.stop instead -->
            <!-- As `row.showDetails` is one-way, we call the toggleDetails function on @change -->
            <b-form-checkbox @click.native.stop @change="row.toggleDetails" v-model="row.detailsShowing">
              Details via check
            </b-form-checkbox>
          </template>
          <template slot="row-details" slot-scope="row">
            <b-card>
              <b-table id="products"
                       striped
                       show-empty
                       :items="row.item.items"
                       :fields="fields_2"
                       ref="table"
              ></b-table>
              <b-button size="sm" @click="row.toggleDetails">Hide Details</b-button>
            </b-card>
          </template>
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
    name: 'Users',
    data () {
      return {
        formUrl: 'http://localhost:5001',
        fields: [
          {
            label: 'Visitor',
            key: 'visitorid_string',
          },
          {
            label: 'Items',
            key: 'show_details',
          },
        ],
        fields_2: [
          {
            label:'Item',
            key:'itemid_string',
          },
          {
            label:'Count',
            key:'itemid_count'
          },
        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy: false,
        totalRows:1,
        currentPage:1,
        perPage:15,
        fileProducts:null,
        items:[]
      }
    },
    created() {
      this.getItems()
    },
    methods:{
      getItems(ctx){
        let url = "http://localhost:5001/users";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.items = result.body;
              this.totalRows = this.items.length;
              this.isBusy = false;
              return result.body
            }
          }
          this.isBusy = false;
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
