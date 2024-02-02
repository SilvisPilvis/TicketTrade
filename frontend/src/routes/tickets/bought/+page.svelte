<script>
    import axios from "$lib/axios";
    import { onMount } from "svelte";
    let result, failure;
    onMount(async () => {
        try{
            const res = await axios.get('/bought/tickets');
            // const res = await axios.get('/events');
            result = res.data;
            console.log(result)
        }catch (e){
            console.error("Error:", e.response);
            failure = e.response;
        }
    })
</script>

<main>
{#if result != "" && result != undefined}
    {#each result as ticket}
        <div>
            <p>{ticket.EventName}</p>
            <p>{ticket.TicketBoughtAt}</p>
            <p>{ticket.UserName}</p>
        </div>
    {/each}
{/if}
</main>