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

<main class="flex col cen">
{#if result != "" && result != undefined}
    <div class="margin-t flex row cen">
        {#each result as ticket}
            <div class="ticket">
                <p>{ticket.EventName}</p>
                <p>{ticket.TicketBoughtAt}</p>
                <p>{ticket.UserName}</p>
            </div>
        {/each}
    </div>
{/if}
</main>

<style>
    .row{
        flex-direction: row;
        flex-wrap: wrap;
    }
    .margin-t{
        margin-top: 5rem;
    }
    .ticket{
        background-color: var(--bg);
        border-radius: 0.4rem;
        margin: 1rem;
        padding: 1rem;
    }
</style>