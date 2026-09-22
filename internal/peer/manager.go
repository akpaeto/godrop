package peer

import "fmt"

type PeerManager struct {
	Peers []*Peer
}

func (pm *PeerManager) AddPeer(newPeer *Peer) {
	for i, existingPeer := range pm.Peers {
		if existingPeer.Name == newPeer.Name {
			pm.Peers[i].Address = newPeer.Address
			fmt.Println("peer уже сущетсвует, обновляем:", existingPeer.Name)

			return
		}
	}

	pm.Peers = append(pm.Peers, newPeer)
	fmt.Println("добавляем новый пир:", newPeer.Name)
}

func (pm *PeerManager) Find(name string) *Peer {
	for _, peer := range pm.Peers {
		if peer.Name == name {
			return peer
		}
	}
	return nil

}

func (pm *PeerManager) List() []*Peer {
	return pm.Peers
}

func (pm *PeerManager) RemovePeer(name string) {
	for i, peer := range pm.Peers {
		if peer.Name == name {
			pm.Peers = append(pm.Peers[:i], pm.Peers[i+1:]...)

			fmt.Println("peer удален:", name)
			return
		}
	}

	fmt.Println("peer для удаления не найден:", name)
}
