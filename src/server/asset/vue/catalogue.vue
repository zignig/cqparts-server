<template>
    <div class="cat" is="sui-card-group">
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    <card></card>
    </div>
</template>

<script>
export default {
    props: {
        catalogue: {
            type: Array,
        },
    },
    methods:{
    }
}
</script>

<style>

.cat {
    float:right;
    }

</style>
