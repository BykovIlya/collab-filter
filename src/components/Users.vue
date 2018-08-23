<template>
    <b-container fluid>
        <b-card border-variant="">
            <template slot="header">
                <b-row>
                    <b-col sm="5">
                        <h4 class="card-title">Table of users</h4>
                    </b-col>
                    <b-col sm="7">
                        <div class ="float-right">
                            <b-btn  variant="info" @click="showModalImport">import from .csv</b-btn>
                        </div>
                    </b-col>
                </b-row>
            </template>
            <b-table id="products"
                     striped
                     show-empty
                     :busy.sync="isBusy"
                     :fields="fields"
                     :current-page="currentPage"
                     :per-page="perPage"
                     :sort-by.sync="sortBy"
                     :sort-desc.sync="sortDesc"
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
        name: 'Users',
        data () {
            return {
                formUrl: 'users',
                fields: [
                    {
                        key: 'id',
                    },
                    {
                        key: 'name',
                        label: this.$t('name'),
                    },
                ],
                currentPage: 1,
                perPage: 15,
                isBusy: false,
                totalRows:0,
                sortBy: null,
                sortDesc: false,
            }
        },
        methods: {
            importProducts() {
                let formData = new FormData();
                formData.append('file', this.fileProducts);
                let url = this.formUrl + "/import";
                this.$http.post(url, formData, null).then(result => {
                    console.log(result);
                    if (result.status === 200) {
                        this.$refs.table.refresh();
                    }
                }, error => {
                    console.log(error);
                });
            },
            showModalImport() {
                this.$refs.modalImport.show()
            },
            hideModalImport() {
                this.$refs.modalImport.hide()
            }
        }
    }

</script>