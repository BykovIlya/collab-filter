<template>
  <div>
    <b-container fluid>
      <b-card border-variant="">
        <template slot="header">
          <b-row>
            <b-col sm="5">
              <h4 class="card-title">Table of events</h4>
            </b-col>
            <b-col sm="7">
              <div class ="float-right">
                <b-btn  variant="info" @click="showModalImport">import from .csv</b-btn>
              </div>
            </b-col>
          </b-row>
        </template>
        <b-table id="home"
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
    <b-modal id="modal-import"
             ref="modalImport"
             title="import"
             @ok="importProducts"
             ok-title = "upload"
             cancel-title = "cancel"
             centered>
      <b-form-file v-model="fileProducts" class="mt-1"></b-form-file>
    </b-modal>
  </div>
</template>

<script>
  export default {
    name: 'Home',
    data () {
      return {
        formUrl: 'http://localhost:5000',
        fields: [
          {
            key: 'timestamp',
          },
          {
            key: 'visitorid',
          },
          {
            key: 'event',
          },
          {
            key: 'itemid',
          },
          {
            key: 'transactionid',
          },
        ],
        newItem: {
          id:null,
          name:'',
        },
        isBusy:false,
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
        let url = "http://localhost:5001/events";
        this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);

          if (result.status === 200 || result.status === 304 ){
            if(result.body.length > 0) {
              this.items = result.body;
              this.totalRows = this.items.length;
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
      importProducts(){
        let formData = new FormData();
        formData.append('file', this.fileProducts);
        let url = this.formUrl+"/import";
        this.$http.post(url, formData, null).then(result => {
          console.log(result);
          if (result.status === 200) {
            this.$refs.table.refresh();
          }
        },error =>{
          console.log(error);
        });
      },
      showModalImport () {
        this.$refs.modalImport.show()
      },

    },
  }

</script>
