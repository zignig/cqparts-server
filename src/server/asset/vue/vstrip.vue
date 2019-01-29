<template>
    <section>
    <nav class="panel">
        <p class="panel-tabs" v-show="visible">
            <a 
                v-for="section in sections"
                v-bind:class="sectionActive(section)"
                :key="section"
                @click="selectSection(section)"
                >{{ section }}</a>
        </p>
        <div class="" v-show="visible">
            <div class="panel-block" :key="model.name" v-for="model in pages">
                <span class="icon" @click="pin(model)">
                    <i v-if="model.pinned" class="mdi mdi-pin"></i>
                    <i v-else="model.pinned" class="mdi mdi-rotate-45 mdi-pin"></i>
                </span>
                <img v-on:click="load(model)" v-show="model.img" :src="model.img"></img>
                <a v-on:click="load(model)"> 
                {{ model.name }}
                </a>
            </div>
            <div class="panel-block">
                <div  class="buttons has-addons " v-show="pageCount>1" attached="top" size="small" v-show="visible" >
                    <a class="button" @click="prevPage" :disabled="currentPage === 0">
                        <span class="icon"><i class="mdi mdi-chevron-double-left"></i></span>
                    </a>
                    <a class="button" > {{ currentPage + 1 }}/{{ pageCount }} </a>
                    <a class="button" @click="nextPage" :disabled="currentPage >= pageCount -1">
                        <span class="icon"><i class="mdi mdi-chevron-double-right"></i></span>
                    </a>
                </div>
            </div>
        </div>
    </nav>
    </section>
</template>

<script>
export default {
    data(){
        return {
            page: { 'All':0,'Pinned':0,'Fix':0},
            visible: true,
            active: 'Pinned',
            sections: ['Pinned','All','Fix']
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
        sectionActive: function(name){
            if (this.active == name){
                return "is-active"
            }
        },
        selectSection: function(name){
            if (this.visible){
                this.active = name
            }
        },
        viz: function(){
            this.visible = !this.visible;
        },
        nextPage: function(){
            this.page[this.active]++;
        },
        prevPage: function(){
            this.page[this.active]--;
        },
        isPinned: function(obj){
            if (obj.pinned){
                return true
            }
        },
        isActive: function(obj){
            if (this.currenti != null ){
                if (this.current.name  == obj.name){
                    return "green"
                }
                else {
                    return "black"
                }
            }
        },
        pin: function(obj){
            obj.pinned = !obj.pinned
            EventBus.$emit('pin',obj.name);
        },
        load: function(obj){
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        },
        pinned : function(model){
            return this.modelList.filter(function(model){
                return model.pinned === true
            })
        },
        currentData: function(){
            let data = []
            if ( this.active  == 'Pinned' ){
                data = this.pinned()
            }
            if ( this.active  == 'All' ){
                data = this.modelList
            }
            if ( this.active  == 'Fix' ){
                data = this.modelList
            }
            return data
        },
        paginatedData: function(){
            const start = this.page[this.active]* this.size,
                end = start + this.size;
            data = this.currentData()
            return data.slice(start, end);
        }
    },
    computed:{
        currentPage(){
            return this.page[this.active]
        },
        pageCount(){
            let l = this.currentData().length,
                s = this.size;
            return Math.ceil(l/s);
        },
        pages() {
            return this.paginatedData();
        }
    }
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
