components {
  id: "tile"
  component: "/scripts/tile.script"
}
components {
  id: "pfx"
  component: "/assets/tile.particlefx"
  position {
    y: 55.0
  }
}
components {
  id: "explode"
  component: "/assets/explode.particlefx"
  position {
    y: -46.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tile\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"tile\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 51.0168\n"
  "  data: 47.04965\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
