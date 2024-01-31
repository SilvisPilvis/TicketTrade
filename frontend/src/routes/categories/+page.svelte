<script>
import axios from "$lib/axios";
axios.get('/categories')
        .then(function (response) {
            console.log(response.data);
        })
        .catch(function (error) {
            console.log(error);
        });
</script>

<main>
<h1>All categories:</h1>
</main>