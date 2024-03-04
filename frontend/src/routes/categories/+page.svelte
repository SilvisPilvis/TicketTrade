<script>
import axios from "$lib/axios";
import { onMount } from "svelte";
import "iconify-icon"
let data, success, failed;
onMount(async () => {
    try{
        const res = await axios.get('/categories');
        data = res.data;
    }catch (e){
        console.error("Error:", e);
    }
})
function deleteCategory(id){
    axios.delete(`/category/${id}`)
      .then(response => {
        // console.log(response.data);
        success = response.data;
      })
      .catch(error => {
        //   console.error(error.response.data.error);
          failed = error.response.data.error;
      });
      return error;
}

//   console.log(data);
</script>

<main class="flex">
<h1 class="margin-t">All categories:</h1>
{#if data != undefined && data != ""}  
    <div class="flex row">
        {#each data as category}
            <div class="category flex row between cen">
                <div>{category.CategoryName}</div>
                <div class="flex row">
                    <a href="../category/{category.Id}/edit" class="flex cen icon"><iconify-icon icon="material-symbols:edit-square" style="color: var(--button-text)"></iconify-icon></a>
                    <button on:click={() => deleteCategory(0)}>
                        <iconify-icon class="flex cen icon" icon="mdi:trash-can"  style="color: var(--button-text)"></iconify-icon>
                    </button>
                </div>
            </div>
        {/each}
        {#if failed != "" && failed != undefined}
            <p class="error">{failed}</p>
        {/if}
        {#if success != "" && success != undefined}
            <p class="error">{success}</p>
        {/if}
    </div>
{/if}
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
    h1{
        margin-left: 1rem;
        width: 100%;
    }
    .category{
        background-color: var(--fg);
        margin: 1rem;
        padding: 0.7rem;
        max-width: 15rem;
        border-radius: 0.4rem;
    }
    .between{
        justify-content: space-between;
    }
    iconify-icon, a{
        cursor: pointer;
        background-color: var(--button-fill);
    }
    button{
        outline: none;
        background-color: var(--button-fill);
        border: none;
    }
    button, a{
        padding: 0.3rem;
        margin: 0.2rem;
        border-radius: 0.2rem;
    }
</style>