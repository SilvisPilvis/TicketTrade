<script>
    import axios from "$lib/axios"
    let categoryName, success;
    
    function createCategory(){
        axios.post('/category/create', {
            categoryName: categoryName
        })
        .then(function (response) {
            console.log(response.data);
            success = response.data;
        })
        .catch(function (error) {
            console.log(error);
        });
    }
</script>

<main>
<input type="text" name="categoryName" id="" bind:value={categoryName}>
{#if success != undefined && success != ""}
    <p class="success">{success}</p>
{/if}
<button on:click={createCategory}>Create Category</button>
</main>