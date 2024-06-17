<script>
  import axios from "axios";

  let error = "";
  let data = [];
  let isLoading = true;

  async function list() {
    try {
      const response = await axios.get("http://localhost:3000/api/anime/list");
      if (response.data.status === "OK") {
        console.log(response.data);
        data = Array.isArray(response.data.data) ? response.data.data : [];
      } else {
        error = "Sorun var.";
      }
    } catch (error) {
      console.error("Hata oluştu:", error);
      error = "Veri alınırken bir hata oluştu.";
    } finally {
      isLoading = false;
    }
  }

  list();
</script>

<main class="p-8">
  <div class="mb-4 text-red-500">{error}</div>
  
  {#if isLoading}
    <p>Yükleniyor...</p>
  {:else if data.length === 0}
    <p>Gösterilecek veri yok.</p>
  {:else}
    <table
      class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
    >
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-4 py-2">Adı</th>
          <th scope="col" class="px-4 py-2">Açıklama</th>
          <th scope="col" class="px-4 py-2">Score</th>
        </tr>
      </thead>
      <tbody>
        {#each data as datas}
          <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
            >
              <button
                class="text-blue font-bold py-1 px-2 rounded"
                on:click={() => (window.location.href = `/anime/${datas.name}`)}
              >
                {datas.name}
              </button>
            </th>
            <td class="px-6 py-4">
              {datas.description}
            </td>
            <td class="px-6 py-4">
              {datas.rating}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</main>
