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

<main>
<h1>All categories:</h1>
{#if data != undefined && data != ""}    
    {#each data as category}
        <div>
            <div>{category.CategoryName}</div>
            <a href="../category/{category.Id}/edit"><iconify-icon icon="material-symbols:edit-square" style="color: black"></iconify-icon></a>
            <iconify-icon icon="mdi:trash-can"  style="color: black" on:click={() => deleteCategory(0)}></iconify-icon>
        </div>
    {/each}
    {#if failed != "" && failed != undefined}
        <p class="error">{failed}</p>
    {/if}
    {#if success != "" && success != undefined}
        <p class="error">{success}</p>
    {/if}
{/if}
</main>

<style>
    
</style>