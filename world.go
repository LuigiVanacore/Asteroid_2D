package asteroid

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	x, y   int
	actors []Actor
}

func NewWorld(x, y int) *World {
	return &World{x: x, y: y}
}

func (w *World) GetX() int {
	return w.x
}

func (w *World) GetY() int {
	return w.y
}

func (w *World) Add(actor Actor) {
	w.actors = append(w.actors, actor)
}

func (w *World) IsCollide(other Collidable) bool {
	for _, actor := range w.actors {
		collider := actor.(Collidable)
		if other.IsCollide(collider) {
			return true
		}
	}
	return false
}

func (w *World) Update() {
	w.updatePositions()
	w.checkCollisions()
	w.deleteActorsDead()
}

func (w *World) updatePositions() {
	for _, actor := range w.actors {
		actor.Update()
		pos := actor.GetTransform().GetPosition()

		outOfBound := false

		if pos.X < 0 {
			pos.X = float64(w.x)
			outOfBound = true
		} else if pos.X > float64(w.x) {
			pos.X = 0
			outOfBound = true
		}
		if pos.Y < 0 {
			pos.Y = float64(w.y)
			outOfBound = true
		} else if pos.Y > float64(w.y) {
			pos.Y = 0
			outOfBound = true
		}

		if (actor.GetTag() == playerShoot) && outOfBound {
			actor.OnDestroy()
		}

		actor.GetTransform().SetPosition(pos.X, pos.Y)
	}
}

func (w *World) deleteActorsDead() {
	for i := len(w.actors) - 1; i >= 0; i-- {
		if !w.actors[i].IsAlive() {
			w.actors = append(w.actors[:i], w.actors[i+1:]...)
		}
	}
}

func (w *World) Clear() {
	w.actors = nil
}

func (w *World) checkCollisions() {
	for i := 0; i < len(w.actors); i++ {
		for j := i + 1; j < len(w.actors); j++ {
			collider1 := w.actors[i]
			collider2 := w.actors[j]
			if collider1.IsAlive() && collider2.IsAlive() {
				if collider1.IsCollide(collider2) {
					collider1.OnDestroy()
					fmt.Printf("first entity %s is colliding with second entity %s", collider1.ToString(), collider2.ToString())
				}
				if collider2.IsCollide(collider1) {
					collider2.OnDestroy()
					fmt.Printf("second entity %s is colliding with first entity %s \n", collider2.ToString(), collider1.ToString())
				}
			}
		}
	}
}

func (w *World) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	//op.GeoM = e.geoM
	for _, actor := range w.actors {
		actor.Draw(target, op)
	}
}

//const auto end = _entities.end();
//for(auto it_i = _entities.begin(); it_i != end; ++it_i)
//{
//    Entity& entity_i = **it_i;
//    auto it_j = it_i;
//    it_j++;
//    for(; it_j != end;++it_j)
//    {
//        Entity& entity_j = **it_j;
//
//        if(entity_i.isAlive() and entity_i.isCollide(entity_j))
//            entity_i.onDestroy();
//
//        if(entity_j.isAlive() and entity_j.isCollide(entity_i))
//            entity_j.onDestroy();
//    }
//}
//
//for(auto it = _entities.begin(); it != _entities.end();)
//{
//    if(not (*it)->isAlive())
//    {
//        delete *it;
//        it = _entities.erase(it);
//    }
//    else
//        ++it;
//}
//
//_sounds.remove_if([](const std::unique_ptr<sf::Sound>& sound) -> bool {
//        return sound->getStatus() != sf::SoundSource::Status::Playing;
//    });

//
//bool World::isCollide(const & other)
//{
//for(Entity* entity_ptr : _entities)
//if(other.isCollide(*entity_ptr))
//return true;
//return false;
//}
//int World::size()
//{
//return _entities.size() + _entitiesTmp.size();
//}
//
//
//
//const std::list<Entity*> World::getEntities()const
//{
//return _entities;
//}
//
//void World::update(sf::Time deltaTime)
//{
//if(_entitiesTmp.size() > 0)
//_entities.splice(_entities.end(),_entitiesTmp);
//
//for(Entity* entity_ptr : _entities)
//{
//Entity& entity = *entity_ptr;
//
//entity.update(deltaTime);
//
//sf::Vector2f pos = entity.getPosition();
//
//if(pos.x < 0)
//{
//pos.x = _x;
//pos.y = _y - pos.y;
//}
//else if (pos.x > _x)
//{
//pos.x = 0;
//pos.y = _y - pos.y;
//}
//
//if(pos.y < 0)
//pos.y = _y;
//else if(pos.y > _y)
//pos.y = 0;
//
//entity.setPosition(pos);
//}
//
//
//const auto end = _entities.end();
//for(auto it_i = _entities.begin(); it_i != end; ++it_i)
//{
//Entity& entity_i = **it_i;
//auto it_j = it_i;
//it_j++;
//for(; it_j != end;++it_j)
//{
//Entity& entity_j = **it_j;
//
//if(entity_i.isAlive() and entity_i.isCollide(entity_j))
//entity_i.onDestroy();
//
//if(entity_j.isAlive() and entity_j.isCollide(entity_i))
//entity_j.onDestroy();
//}
//}
//
//for(auto it = _entities.begin(); it != _entities.end();)
//{
//if(not (*it)->isAlive())
//{
//delete *it;
//it = _entities.erase(it);
//}
//else
//++it;
//}
//
//_sounds.remove_if([](const std::unique_ptr<sf::Sound>& sound) -> bool {
//return sound->getStatus() != sf::SoundSource::Status::Playing;
//});
//}
//
//
//
//void World::draw(sf::RenderTarget& target, sf::RenderStates states) const
//{
//for(Entity* entity : _entities)
//target.draw(*entity,states);
//}
//}
//
