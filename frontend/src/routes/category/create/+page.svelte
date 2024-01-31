<script>
    import axios from "$lib/axios"
    let categoryName;
    
    function createCategory(){
        axios.post('/category/create', {
            categoryName: categoryName
        })
        .then(function (response) {
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
    }
</script>

<main>
<input type="text" name="categoryName" id="" bind:value={categoryName}>
<button on:click={createCategory}>Create Category</button>
</main>