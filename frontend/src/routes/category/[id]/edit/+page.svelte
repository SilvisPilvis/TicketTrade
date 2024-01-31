<script>
    import axios from "$lib/axios";
    export let data;
    let categoryName;

    console.log(data.data);

    // console.log(page.params.id);

    axios.get(`/category/${data.data}`)
        .then(function (response) {
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
    
    function editCategory(){
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
    Edit Category:
<input type="text" name="categoryName" id="" bind:value={categoryName} placeholder="Category Name">
<button on:click={editCategory}>Create Category</button>
</main>