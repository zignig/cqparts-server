
# run me from blender 
import json
import requests
import time
import os
import zipfile ,io 
import bpy

target =  os.environ['CQPARTS_SERVER']
def tryAgain(retries=0):
    if retries > 50: return
    is_data = False
    try:
        r = requests.get(target+'/render', stream=True)
        for line in r.iter_lines():
            if line:
                decoded_line = line.decode('utf-8')
                spl = decoded_line.split(":",1)
                print(spl)
                if is_data:
                    if spl[0] == 'data':
                        data = spl[1]
                        jdata = json.loads(data)
                        is_data == False
                        render_this(jdata)

                if spl[0] == "event":
                    if spl[1] == 'render':
                       is_data = True

    except Exception as e:
        print(e)
        time.sleep(2)
        retries+=1
        print ("retries :"+str(retries))
        tryAgain(retries)

def render_this(jdata):
    name = jdata['name']
    print(jdata)
    r = requests.get(target+'/zipped/'+name,stream=True)
    z= zipfile.ZipFile(io.BytesIO(r.content))
    z.extractall("/opt/stash/")
    print(r)
    make_blender(name,jdata['cam'],jdata['target'])
    uploader(name)


# using
# https://github.com/ksons/gltf-blender-importer
# snippit , not working script
def make_blender(name,cam_loc,tgt_loc):
    multiplier =  100
    res = (800,600) # pass down from parts server 
    samples =  3 
    size_per = 20
    bpy.ops.wm.addon_enable(module="io_scene_gltf")
    # maybe script build the entire scene
    bpy.ops.scene.new(type='NEW')
    #bpy.context.scene.name = 'cqparts'
    # make the world
    bpy.ops.world.new()
    world = bpy.data.worlds['World.001']
    world.name = 'NewWorld'
    world.light_settings.use_ambient_occlusion = True

    bpy.context.scene.world = world

    #theScene = bpy.data.scenes['cqparts']
    theScene = bpy.context.scene
    theScene.cycles.samples = samples 
    theScene.render.filepath = "/opt/stash/"+name+".png"
    theScene.render.resolution_x = res[0] 
    theScene.render.resolution_y = res[1] 
    theScene.render.resolution_percentage = size_per 
    # make and bind the camera
    bpy.ops.object.camera_add()
    cam = bpy.context.selected_objects[0]
    bpy.context.scene.camera = cam
    cam.location = (-cam_loc['x']*multiplier,-cam_loc['y']*multiplier,-cam_loc['z']*multiplier)
    # add the track
    bpy.ops.object.constraint_add(type="TRACK_TO")

    # make the target and track the camera
    bpy.ops.object.empty_add(type='SPHERE')
    tgt  = bpy.context.selected_objects[0]
    tgt.name = "cam_target"
    tgt.location = (-tgt_loc['x']*multiplier,-tgt_loc['y']*multiplier,-tgt_loc['z']*multiplier)
    # select the camers
    track = cam.constraints["Track To"]
    track.target = bpy.data.objects['cam_target']
    track.up_axis = 'UP_Y'
    track.track_axis = 'TRACK_NEGATIVE_Z'


    # lamp 1
    bpy.ops.object.lamp_add(type='POINT')
    lamp = bpy.context.selected_objects[0]
    lamp.location = (10,-10,10)

    # lamp 2
    bpy.ops.object.lamp_add(type='POINT')
    lamp2 = bpy.context.selected_objects[0]
    lamp2.location = (-10,-10,10)

    bpy.ops.import_scene.gltf(filepath="/opt/stash/"+name+"/out.gltf")
    # this does not work second time.
    try:
        outer = theScene.objects['out']
        outer.scale = (100,100,100)
        bpy.ops.render.render(write_still=True)
        outer.select
        bpy.ops.object.delete()
        bpy.ops.scene.delete()
    except:
        for i in theScene.objects:
            print(i)

    for i in theScene.objects:

        i.select = True
    bpy.ops.object.delete()
    print(theScene.objects)

def uploader(name):
    file_ref = ('objs',(name+"png",open("/opt/stash/"+name+".png","rb")))
    requests.post(target+'/image',files=[file_ref])

print ("Running Blender Render runner")
tryAgain()
