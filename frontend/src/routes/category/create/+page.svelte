<script>
    import axios from "$lib/axios"
    let categoryName, success, fail;
    
    function createCategory(){
        console.log(categoryName);
        if(categoryName === undefined){
            fail = "Category name can't be empty.";
            return;
        }else{
            fail = "";
        }
        if(categoryName != "" || categoryName !== undefined){
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
        }else{
            fail = "Category name can't be empty.";
        }
        // console.log(categoryName.length);
    }
</script>

<main class="flex col cen">
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
    .margin-t{
        margin-top: 5rem;
    }
    main{
        height: 100%;
    }
    input {
        /* align-self: flex-start; */
        padding: 1rem;
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: var(--fg);
        font-size: 12pt;
        margin: 1rem 1rem 1rem 0rem;
    }
    button {
        border-radius: 0.4rem;
        padding: 1.1rem 3.7rem 1.1rem 3.7rem;
        border: none;
        outline: none;
        background-color: var(--button-fill);
    }
    label{
        margin: 1rem 0 0 1rem;
    }
</style>