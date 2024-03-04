<script>
import axios from "$lib/axios";
import { onMount } from "svelte";
import "iconify-icon"

let res, success, fail;
onMount(async () => {
    try{
        //get event data
        const response = await axios.get('/ticket/types');
        res = response.data;
        // console.log(res);
    }catch (e){
        console.error("Error:", e);
    }
})

function deleteTicketType(id){
    axios.delete(`/ticket/type/${id}`)
            .then(function (response) {
                console.log(response.data);
                success = response.data;
                fail = "";
            })
            .catch(function (error) {
                console.log(error);
                fail = error.response.data.error;
                success = "";
            });
}
</script>

<main class="flex col cen">
{#if res != "" && res != undefined}
    <div class="margin-t flex row cen">
        {#each res as type}
            <div class="flex row cen type">
                <p>{type.typeName}</p>
                <p> </p>
                <p>{type.typePrice} €</p>
                <a href={"/ticket/types/"+type.Id+"/edit"} class="flex cen"><iconify-icon icon="material-symbols:edit"  style="background-color: var(--button-fill);"></iconify-icon></a>
                <button on:click={() => deleteTicketType(type.Id)}><iconify-icon icon="mdi:trash-can"  style="background-color: var(--button-fill);"></iconify-icon></button>
            </div>
        {/each}
        {#if fail != "" && fail != undefined}
            <p class="error">{fail}</p>
        {/if}
        {#if success != "" && success != undefined}
            <p class="success">{success}</p>
        {/if}
    </div>
{/if}
</main>

<style>
    .type{
        background-color: var(--fg);
        margin: 0.2rem;
        border-radius: 0.4rem;
        padding: 0.4rem;
    }
    main{
        width: 100%;
        height: 100%;
    }
    .margin-t{
        margin-top: 5rem;
    }
    button, a{
        width: 1.7rem;
        height: 1.7rem;
        border: none;
        margin: 0.5rem;
        padding: 0.2rem;
        background-color: var(--button-fill);
        border-radius: 0.4rem;
    }
</style>