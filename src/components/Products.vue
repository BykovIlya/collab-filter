<template>
    <b-container fluid>
        <b-card border-variant="">
            <template slot="header">
                <b-row>
                    <b-col sm="5">
                        <h4 class="card-title">Table of products</h4>
                    </b-col>
                    <b-col sm="7">
                        <div class ="float-right">
                            <b-btn  variant="info" @click="showModalImport">import from .csv</b-btn>
                            <b-btn  variant="success" @click="showModal">import 2</b-btn>
                        </div>
                    </b-col>
                </b-row>
            </template>
            <b-table id="products"
                     striped
                     show-empty
                     :busy.sync="isBusy"
                     :items="getItems"
                     :fields="fields"
                     :current-page="currentPage"
                     :per-page="perPage"
                     ref="table"
            >
    </b-table>
    <b-row>
        <b-col sm="12">
            <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage" class="my-0" />
        </b-col>
    </b-row>
        </b-card>
    </b-container>

</template>

<script>
    export default {
        name: 'Products',
        data () {
            return {
                formUrl: 'products',
                fields: [
                    {
                        key: 'id',
                    },
                    {
                        key: 'name',
                        label: this.$t('name'),
                    },
                ],
                newItem: {
                    id:null,
                    name:'',
                },
                isValidBarcode:true,
                currentPage: 1,
                perPage: 15,
                isBusy: false,
                totalRows:0,
                deleteRow:{
                    id:null,
                    name:'',
                },
                fileProducts:null
            }
        },
        methods:{
            getItems(ctx){
                let url = this.formUrl+"?page="+this.currentPage+"&size="+this.perPage;
                this.isBusy = true;
                return this.$http.get(url).then(result => {
                    console.log(result);
                    this.isBusy = false;
                    if (result.status === 403) {
                        this.showForbiddenAlert = true;
                    }else if (result.status === 200 || result.status === 304 ){
                        this.showForbiddenAlert = false;
                        if(result.body.products.length > 0) {
                            this.totalRows = result.body.total;
                            return result.body.products
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
            clearNewItem(){
                this.newItem.id = null;
                this.newItem.name = '';
                this.newItem.barcode = '';
                this.newItem.weight = 0;
                this.newItem.price = 0;
                this.newItem.sale = 0;
            },
            createProduct (evt) {
                evt.preventDefault()
                this.handleSubmit()
            },
            handleSubmit () {
                var url = this.formUrl
                this.isValidBarcode = true
                if (this.nameValid() && this.saleValid() && this.priceValid()){
                    if (this.newItem.id > 0) {
                        url = this.formUrl+'/'+this.newItem.id
                        this.put(url,this.newItem,function(result){
                            console.log(result.body)
                            if (result.status === 422){
                                this.isValidBarcode = false
                            }else if (result.status === 200){
                                this.$refs.table.refresh();
                                this.hideModal()
                            }}.bind(this));
                    }else{
                        this.post(url,this.newItem,function(result){
                            console.log(result)
                            if (result.status === 422){
                                this.isValidBarcode = false
                            }else  if (result.status === 200) {
                                this.$refs.table.refresh();
                                this.hideModal()
                            }}.bind(this));
                    }
                }
            },
            removeItem(row){
                let url = this.formUrl+'/'+row.id
                this.delete(url,null,function(result){
                    if (result.status === 422){
                        //this.formErrors = result.errors
                    }else {
                        this.$refs.table.refresh();
                        this.hideModalDelete()
                    }}.bind(this));
            },
            showModal (row) {
                if (row.item){
                    this.newItem.id = row.item.id
                    this.newItem.name = row.item.name
                }else{
                    this.clearNewItem()
                }
                this.$refs.modal.show()
            },
            hideModal(){
                this.$refs.modal.hide()
            },
            showModalDelete (row) {
                if (row.item){
                    this.deleteRow =row.item
                }
                this.$refs.modalDelete.show()
            },
            hideModalDelete(){
                this.$refs.modalDelete.hide()
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
            hideModalImport(){
                this.$refs.modalImport.hide()
            },
        },
        computed: {
            validateNameState () {
                return this.nameValid()
            },
            validateSaleState () {
                return this.saleValid()
            },
            validatePriceState () {
                return this.priceValid()
            },
            validateBarcodeState () {
                return this.isValidBarcode
            },
            validateWeightState () {
                return this.weightValid()
            }
        },
        watch: {
            'newItem.barcode':{
                handler:function (newVal, oldVal) {
                    if (!this.isValidBarcode) {
                        this.isValidBarcode = newVal !== oldVal
                    }
                },
                deep: true
            }
        }
    }

</script>



