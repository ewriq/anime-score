<script>
  import axios from "axios";
  import Cookies from "js-cookie";
  import Del from "./Del.svelte";
  let user = Cookies.get("token");

  let error = "";
  let data = [];
  async function list() {
    try {
      const response = await axios.post("http://localhost:3000/api/vote/list", {
        token: user,
      });
      if (response.data.status === "OK") {
        console.log(response.data);
        data = response.data.data;
      } else {
        if (
          response.data.error &&
          Object.keys(response.data.error).length === 0
        ) {
          error = "Data yok";
        } else {
          error = "Hata var";
        }
      }
    } catch (error) {
      console.error("Hata oluştu:", error);
    }
  }
  list();
</script>

<main class="p-8">
  {#if error == "Hata var"}
    <div class="mb-4 text-red-500">{error}</div>
  {:else if error == ""}
  <h2 class="px-6 py-4 whitespace-nowrap dark:text-white">Kullandıgın oylar tablosu </h2>
    <table
      class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"
    >
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-4 py-2">Adı</th>
          <th scope="col" class="px-4 py-2">Verdigin Oy </th>
          <th scope="col" class="px-4 py-2">Oyu sil</th>
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
                class=" text-blue font-bold py-1 px-2 rounded"
                on:click={() =>
                  (window.location.href = `/anime/${datas.anime}`)}
              >
                {datas.anime}
              </button>
            </th>
            <td class="px-6 py-4">
              {datas.score}
            </td>

            <td class="px-6 py-4">
              <Del
                anime={datas.anime}
                score={datas.score}
                voter={datas.voter}
              />
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else if error == "Data yok"}
    <div>
    </div>
  {/if}
</main>
