// A viewer for cqparts
//
var container, controls , obj;
var camera, scene, renderer, light;
var meshlist = [];
var raycaster ;
var mouse = new THREE.Vector2(), INTERSECTED ; 

init();
animate();

function clear(){
    scene.remove(obj);
}

function load(name){
	var loader = new THREE.GLTFLoader();
    loader.load( 'model/'+name+'/out.gltf', function ( gltf ) {
		scene.add( gltf.scene );
        obj = gltf.scene;
		//grab all the meshes for selection
        meshlist = [];
        gltf.scene.traverse( function ( child ) {
            if ( child.isMesh ) {
                meshlist.push(child);
            }
            } );
        } );
    }

function init() {
	canvas = document.getElementById('viewer');

    camera = new THREE.PerspectiveCamera( 30, canvas.clientWidth/ canvas.clientHeight, 0.001, 1000);
    camera.fov = 35;
	camera.position.set( 0.2, 0.2, 0.2);
	controls = new THREE.OrbitControls( camera , canvas );
    //controls.autoRotate = true;
    controls.autoRotateSpeed = 2;
	controls.target.set(0,0,0);
	controls.update();
    scene = new THREE.Scene({antialias: true});
    scene.background =  new THREE.Color(255,255,255)
    // light 1
	light = new THREE.HemisphereLight( 0xbbbbff, 0x444422 );
	light.position.set( 0, 20, 0 );
    scene.add( light );
    // light 2
	light2 = new THREE.PointLight(0xf0f0f0,2,100);
	light2.position.set( 50,50,50);
	scene.add( light2 );
    // light 3
	light3 = new THREE.PointLight(0xf0f0f0,2,100);
	light3.position.set( 0,0,-50);
	scene.add( light2) ;
    // renderer
	renderer = new THREE.WebGLRenderer( { canvas: canvas, antialias: true, preserveDrawingBuffer: true } );
	renderer.setPixelRatio( window.devicePixelRatio );
    renderer.setViewport(0,0,window.innerWidth,window.innerHeight);
	renderer.gammaOutput = true;
    // raycaster
    raycaster = new THREE.Raycaster();
    // add to doc and bind events
    window.addEventListener( 'resize', onWindowResize, false );
    canvas.addEventListener('mousemove',onDocumentMouseMove,false);
    canvas.addEventListener('mousedown',onDocumentClick,false);
    document.addEventListener('keydown',onKey,false);
    // Base grid helps us orient ourselves
    var baseGrid = new THREE.GridHelper(1, 10);
    scene.add(baseGrid);
    var baseGridSmall = new THREE.GridHelper(0.1, 10);
    scene.add(baseGridSmall);
    // final resize
    onWindowResize();
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
    requestAnimationFrame( animate );
    controls.update();
	render();
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
