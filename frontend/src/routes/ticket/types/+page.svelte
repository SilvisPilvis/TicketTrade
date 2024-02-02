<script>
import axios from "$lib/axios";
import { onMount } from "svelte";
import "iconify-icon"

let res;
onMount(async () => {
    try{
        //get event data
        const response = await axios.get('/ticket/types');
        res = response.data;
        console.log(res);
    }catch (e){
        console.error("Error:", e);
    }
})

// axios.get('/ticket/types')
//         .then(function (response) {
//             console.log(response.data);
//         })
//         .catch(function (error) {
//             console.log(error);
//         });
</script>

<main>
{#if res != "" && res != undefined}
    {#each res as type}
        <div>
            <p>{type.TicketTypeName}</p>
            <p>{type.TicketTypePrice} €</p>
            <!-- <p>{type.TicketTypePrice} <iconify-icon icon="material-symbols:euro"  style="color: black"></iconify-icon></p> -->
        </div>
    {/each}
{/if}
</main>