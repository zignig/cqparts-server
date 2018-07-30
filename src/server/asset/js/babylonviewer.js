// A viewer for cqparts
//

var canvas , engine , scene , camera , thing; 
init();
animate();

function clear(){
}

function load(name){
}

function init() {
    canvas = document.getElementById('viewer');
    canvas.oncontextmenu = function () { return false; }
    engine = new BABYLON.Engine(canvas, true);
    scene = new BABYLON.Scene(engine);
    camera = new BABYLON.ArcRotateCamera('camera',0,0,0, new BABYLON.Vector3(0, 0.6,-0.6), scene);
    camera.minZ = 0;


    thing = BABYLON.SceneLoader.Append("model/Rover/", "out.gltf", scene, function (scene) {
        var light = new BABYLON.DirectionalLight("dir01", new BABYLON.Vector3(-1, -2, -1), scene);
        light.position = new BABYLON.Vector3(-2, 10, -1);
        light.setDirectionToTarget(new BABYLON.Vector3(0, 1, 0));
        light.intensity = 1;
        light.shadowMinZ = 1;
        light.shadowMaxZ = 10;
        light.shadowEnabled = true;

        var light2 = new BABYLON.DirectionalLight("dir02", new BABYLON.Vector3(-1, -2, -1), scene);
        light2.position = new BABYLON.Vector3(2, 4, -1);
        light2.setDirectionToTarget(new BABYLON.Vector3(0, 0.5, 0));
        light2.intensity = 1;
        light2.shadowMinZ = 1;
        light2.shadowMaxZ = 10;
        light2.shadowEnabled = true;

        var generator = new BABYLON.ShadowGenerator(512, light2);
        generator.useBlurCloseExponentialShadowMap = true;
        generator.useKernelBlur = true;
        generator.blurScale = 1.0;
        generator.blurKernel = 10.0;

        for (var index = 0; index < scene.meshes.length; index++) {
            var m = scene.meshes[index];
            if (m === sky) {
                continue;
            }
            m.receiveShadows = true;
            generator.getShadowMap().renderList.push(m);
        }

        scene.materials[0].environmmentIntensity = 0.1;

    });
    camera.setTarget(BABYLON.Vector3.Zero());
    camera.attachControl(canvas, false);
    //    scene.createDefaultCameraOrLight(true, true, true);
    cam  = scene.activeCamera
    cam.wheelPrecision = 300.0;
    cam.inertia = 0.7;
    cam.panningInertia = 0.1;

    engine.runRenderLoop(function(){
        scene.render();
    });
    window.addEventListener('resize', function() {
        engine.resize();
    });
    return scene;
}

function onKey( event ) {
    if (event.code == "KeyA"){
        //console.log(event);
        for ( var i in meshlist){
            meshlist[i].visible = true;
        }
    };
    //console.log(event);
    if (event.code == "Space"){
        if ( INTERSECTED ) {
            INTERSECTED.visible = false;
        };
    };
}

function onDocumentClick( event ) {
    if ( INTERSECTED ) {
        //console.log(INTERSECTED);
        EventBus.$emit("click",INTERSECTED.name);
        //INTERSECTED.visible = false;
    };
}

function onDocumentMouseMove( event ) {
    event.preventDefault();
    var rect = renderer.domElement.getBoundingClientRect();
    mouse.x = ( ( event.clientX - rect.left ) / (rect.right - rect.left)) * 2 - 1;
    mouse.y = - ( ( event.clientY - rect.top ) / (rect.bottom - rect.top)) * 2 + 1;
}

function onWindowResize() {
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize( window.innerWidth,window.innerHeight);
}

function animate() {
    //requestAnimationFrame( animate );
    //controls.update();
    //render();
}

function render(){
    raycaster.setFromCamera(mouse,camera);
    var intersects = raycaster.intersectObjects(meshlist);
    if (intersects.length > 0){
        if ( INTERSECTED != intersects[0].object){
            if (INTERSECTED) INTERSECTED.material.emissive.setHex(INTERSECTED.currentHEX);
            INTERSECTED = intersects[0].object;
            INTERSECTED.currentHex = INTERSECTED.material.emissive.getHex();
            INTERSECTED.material.emissive.setHex(0x252500);
        }
    } else {
        if ( INTERSECTED ) INTERSECTED.material.emissive.setHex(INTERSECTED.currentHex);
        INTERSECTED = null;
    }
    renderer.render( scene, camera );
}
