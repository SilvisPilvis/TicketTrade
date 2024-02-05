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
        <div class="category flex row between cen">
            <div>{category.CategoryName}</div>
            <div class="flex row">
                <a href="../category/{category.Id}/edit"><iconify-icon icon="material-symbols:edit-square" style="color: black"></iconify-icon></a>
                <iconify-icon class="flex cen" icon="mdi:trash-can"  style="color: black" on:click={() => deleteCategory(0)}></iconify-icon>
            </div>
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
    h1{
        margin-left: 1rem;
    }
    .category{
        background-color: var(--bg);
        margin: 1rem;
        padding: 0.7rem;
        max-width: 15rem;
        border-radius: 0.4rem;
    }
    .between{
        justify-content: space-between;
    }
    iconify-icon{
        cursor: pointer;
    }
    iconify-icon, a{
        padding: 0.3rem;
    }
</style>