<template>
<form id="uploadForm" name="uploadForm" enctype="multipart/form-data">

  <input type="file" id="files" name="files" multiple><br>
  <input type="text" id="name" name="name"><br>
  <input type="email" id="email" name="email">


  <input type=button value=Upload @click="this.uploadFiles">

</form>
</template>
<script>
  import axios from 'axios'
  export default {
    name: 'hello',
    data () {
      return {
      }
    },
    methods: {
      uploadFiles () {
        var s = this
        const data = new FormData(document.getElementById('uploadForm'))
        var imagefile = document.querySelector('#file')
        console.log(imagefile.files[0])
        data.append('file', imagefile.files[0])
        data.append('name', s.name)
        data.append('email', s.email)
        axios.post('http://192.168.1.222:8080/upload', data, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
          .then(response => {
            console.log(response)
          })
          .catch(error => {
            console.log(error.response)
          })
      }
    }
  }
</script>
