<script>
    import axios from "$lib/axios";
    // import { onMount } from "svelte";
    export let data;
    let categoryName, response, success, fail;
    
    axios.get(`/category/${data.data}`)
        .then(function (response) {
            console.log(response);
            response = response;
            categoryName = response.data.CategoryName;
        })
        .catch(function (error) {
            console.log(error);
        });
    // console.log(response);

    
    function editCategory(){
        if(data.data <= 0){
            fail = "The category Id mmust me greater than 0.";
        }
        axios.put(`/category/${data.data}`, {
        // axios.put(`/category/${-1}`, {
            categoryName: categoryName
        })
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            fail = error.response.data.error;
        });
    }
    // console.log(success);
</script>

<main>
    Edit Category:
<input type="text" name="categoryName" id="" bind:value={categoryName} placeholder="Category Name">
{#if fail != undefined && fail != ""}
    <p class="error">{fail}</p>
{/if}
{#if success != undefined && success != ""}
    <p class="success">{success}</p>
{/if}
<button on:click={editCategory}>Edit Category</button>
</main>