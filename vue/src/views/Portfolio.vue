<template>
  <div class="container">
    <div class="row row-sm-revers">
      <div class="col-md-6">
        <SpAssets />
        <SpAmountSelect />
      </div>
      <div class="col-md-5 col-lg-4 col-md-offset-1 col-lg-offset-2">
        <SpAcc />
      </div>
    </div>
  </div>
</template>

<script>
import { SpAssets, SpAcc, SpAmountSelect } from '@starport/vue'
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'Portfolio',

  components: { SpAcc, SpAssets, SpAmountSelect },

  setup() {
    // store
    let $s = useStore()

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    return {
      address
    }
  }
}
</script>

<style scoped>
.row {
  display: flex;
  flex-wrap: wrap;
}
.col {
  flex-grow: 1;
  padding: 20px;
}
</style>
