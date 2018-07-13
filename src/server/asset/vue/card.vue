<template lang="html">
  <div>
      <sui-card>
        <sui-card-content>
          <sui-card-header>Rover</sui-card-header>
          <sui-card-description>
            Awesome rover thing 
          </sui-card-description>
        </sui-card-content>
        <sui-image src="/pic/Rover.png"></sui-image>
        <sui-button-group>
            <sui-button >
            <sui-icon name="truck"></sui-icon>
               Add 
            </sui-button>
            <sui-button >
            <sui-icon name="eye"></sui-icon>
               View 
            </sui-button>
        </sui-button-group>
      </sui-card>
  </div>
</template>

<script>
export default {
    props : {}
};
</script>
