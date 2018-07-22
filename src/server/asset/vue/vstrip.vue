<template>
    <sui-segment>
        <sui-button @click="viz" :positive="visible" basic compact icon="bars"></sui-button>
        <sui-button @click="viz" :positive="visible" circular basic compact icon="info circle"></sui-button>
        <sui-item-group v-show="visible" divided >
            <sui-button-group attached="top" size="small" v-show="visible" >
                <sui-button @click="prevPage" :disabled="pageNumber === 0" icon="caret left"></sui-button>
                <sui-button > {{ pageNumber + 1 }}/{{ pageCount }} </sui-button>
                <sui-button @click="nextPage" :disabled="pageNumber >= pageCount -1" icon="caret right" floated="right"></sui-button>
            </sui-button-group>
            <sui-item :key="model.name" v-for="model in paginatedData">
                <sui-button class="nogap" basic v-on:click="load(model)">
                    <sui-item-image wrapped v-on:click="load(model)" size="tiny" v-show="model.img" :src="model.img">
                    </sui-item-image>
                </sui-button>
                <sui-item-content>
                    <h4 is="sui-header" :color="isActive(model)">
                        <span><sui-icon size="small" name="delete" v-on:click="remove(model)"/></sui-icon></span>
                        {{ model.name }}
                    </h4>
                    <sui-item-extra>
                        <sui-button-group>
                            <sui-button size="mini" icon="camera" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="newspaper" v-on:click=""></sui-button>
                            <sui-button size="mini" icon="thumbtack" v-on:click=""></sui-button>
                        </sui-button-group>
                    </sui-item-extra>
                </sui-item-content>
            </sui-item>
            <sui-button-group attached="bottom" size="small" v-show="visible" >
                <sui-button @click="prevPage" :disabled="pageNumber === 0" icon="caret left"></sui-button>
                <sui-button > {{ pageNumber + 1 }}/{{ pageCount }} </sui-button>
                <sui-button @click="nextPage" :disabled="pageNumber >= pageCount -1" icon="caret right" floated="right"></sui-button>
            </sui-button-group>
        </sui-item-group>
    </sui-segment>
</template>

<script>
export default {
    data(){
        return {
            pageNumber: 0,
            visible: true 
        }
    },
    props: {
        current: "",
        modelList: {
            type: Array,
        },
        size:{
            type: Number,
            required: false,
            default: 6,
        }
    },
    methods:{
        viz: function(){
            this.visible = !this.visible;
        },
        nextPage: function(){
            this.pageNumber++;
        },
        prevPage: function(){
            this.pageNumber--;
        },
        isActive: function(obj){
            if (this.current == obj.name){
                return "green"
            }
            else {
                return "black"
            }
        },
        load: function(obj){
            clear();
            load(obj.name);
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        }
    },
    computed:{
        pageCount(){
            let l = this.modelList.length,
                s = this.size;
            return Math.ceil(l/s);
        },
        paginatedData(){
            const start = this.pageNumber * this.size,
                end = start + this.size;
            return this.modelList
                .slice(start, end);
        }}
};

</script>

<style>
.nogap {
    padding-top: 0.0 !important;
    padding-right: 0.0 !important;
    padding-left: 0.0 !important;
    padding-bottom: 0.0 !important;
}
</style>
