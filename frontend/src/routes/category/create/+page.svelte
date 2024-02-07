<script>
    import axios from "$lib/axios"
    let categoryName, success, fail;
    
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
            fail = error.response.data.error;
            success = "";
        });
    }
</script>

<main class="flex col">
<label class="flex col">
    Category name:
    <input type="text" name="categoryName" id="" bind:value={categoryName}>
</label>
{#if success != undefined && success != ""}
    <p class="success">{success}</p>
{/if}
{#if fail}
    <p class="error">{fail}</p>
{/if}
<button on:click={createCategory}>Create Category</button>
</main>

<style>
 input {
        align-self: flex-start;
        padding: 1rem;
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: var(--bg);
        font-size: 12pt;
        margin: 1rem 1rem 1rem 0rem;
    }
    button {
        border-radius: 0.4rem;
        margin-left: 1rem;
        align-self: flex-start;
        padding: 1.1rem 3.5rem 1.1rem 3.25rem;
        border: none;
        outline: none;
    }
    label{
        margin: 1rem 0 0 1rem;
    }
</style>