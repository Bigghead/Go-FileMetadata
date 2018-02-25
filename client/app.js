
const onFileChange = (event) => {

    let headers;
    let fileList = event.files;
    if(fileList.length > 0) {
        let file = fileList[0];
        let formData = new FormData();
        formData.append('uploadFile', file, file.name);

        fetch( 'http://localhost:9000/file', {
            method: 'POST',
            body: formData
        } )
        .then( res => console.log( 'Success' ) )
        .catch( err => console.log( 'Son, I am dissapoint' ) )

    }
  }

document.addEventListener( 'DOMContentLoaded', () => {

    
})